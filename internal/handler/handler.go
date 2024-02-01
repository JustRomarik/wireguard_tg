package handler

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/romus204/wireguard_tg/internal/config"
	"github.com/skip2/go-qrcode"
	tg "gopkg.in/telebot.v3"
)

var cfg = config.GetConfig()

func Echo(c tg.Context) error {

	return c.Send(c.Text())
}

func ServerON(c tg.Context) error {

	cmd := exec.Command("systemctl", "start", cfg.WgServiceName)

	if err := cmd.Run(); err != nil {
		return c.Send(err)
	} else {
		return c.Send("Server ON")
	}
}

func ServerOFF(c tg.Context) error {

	cmd := exec.Command("systemctl", "stop", cfg.WgServiceName)

	if err := cmd.Run(); err != nil {
		return c.Send(err.Error())
	} else {
		return c.Send("Server OFF")
	}
}

func GetConfig(c tg.Context) error {

	f, err := os.ReadFile(cfg.WgConfPath)

	if err != nil {
		log.Println(err)
		return c.Send(err.Error())
	}

	return c.Send(string(f))
}

func AllText(c tg.Context) error {

	messageText := c.Text()

	if strings.HasPrefix(messageText, "/newuser") {
		if data := strings.Split(messageText, " ")[1:]; len(data) != 0 {

			username := data[0]
			allowedIPs := data[1]
			privateKey := generatePrivateKey()
			publicKey := generatePublicKey(privateKey)

			if err := writeUserToWgConfig(username, allowedIPs, publicKey); err != nil {
				return c.Send(err.Error())
			}

			qrStr, err := generateUserConfig(username, privateKey, allowedIPs)

			if err != nil {
				return c.Send(err.Error())
			}

			if err := sendqr(c, qrStr); err != nil {
				return c.Send(err.Error())
			}

			if err := restartWgService(); err != nil {
				return c.Send(err.Error())
			}
		}
		return c.Send("Good!")
	}

	return c.Send("Try again")

}

func generatePrivateKey() string {

	buf := bytes.Buffer{}

	cmd := exec.Command("wg", "genkey")

	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	return buf.String()[:44]

}

func generatePublicKey(privateKey string) string {

	buf := bytes.Buffer{}

	cmd := exec.Command("wg", "pubkey")
	cmd.Stdin = bytes.NewReader([]byte(privateKey))

	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	return buf.String()[:44]

}

func writeUserToWgConfig(username, allowedIPs, publicKey string) error {

	f, err := os.OpenFile(cfg.WgConfPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("\n[Peer] # %s \nPublicKey = %s \nAllowedIPs = %s\n", username, publicKey, allowedIPs))

	if err != nil {
		log.Println(err)
	}

	return nil

}

func generateUserConfig(username, privateKey, allowedIPs string) (string, error) {

	confg_layout := `[Interface] # %s
					PrivateKey = %s
					Address = %s
					DNS = 8.8.8.8
					
					[Peer]
					PublicKey = %s
					Endpoint = %s
					AllowedIPs = 0.0.0.0/0
					PersistentKeepalive = 20`

	serverAddress := fmt.Sprintf("%s:%s", cfg.ServerAddress, cfg.WgPort)
	conf := fmt.Sprintf(confg_layout, username, privateKey, allowedIPs, cfg.WgPubKey, serverAddress)

	return conf, nil

}

func sendqr(c tg.Context, qrStr string) error {

	qr, err := qrcode.Encode(qrStr, qrcode.Highest, 512)

	if err != nil {
		return err
	}

	f, err := os.CreateTemp("", "*.png")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	defer os.Remove(f.Name())

	if _, err := f.Write(qr); err != nil {
		log.Fatal(err)
	}

	name := f.Name()

	photo := tg.Photo{File: tg.FromDisk(name)}

	return c.Send(&photo)

}

func restartWgService() error {

	cmd := exec.Command("systemctl", "restart", cfg.WgServiceName)

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil

}
