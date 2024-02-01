package main

import (
	"log"
	"time"

	"github.com/romus204/wireguard_tg/internal/config"
	h "github.com/romus204/wireguard_tg/internal/handler"
	tg "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func main() {

	cfg := config.GetConfig()

	pref := tg.Settings{
		Token:  cfg.TgToken,
		Poller: &tg.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tg.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Use(middleware.Whitelist(cfg.TgIdAllowed))

	bot.Handle("/echo", h.Echo)
	bot.Handle("/serveron", h.ServerON)
	bot.Handle("/serveroff", h.ServerOFF)
	bot.Handle("/getconfig", h.GetConfig)
	bot.Handle(tg.OnText, h.AllText)

	bot.Start()
}
