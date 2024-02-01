package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var once sync.Once
var Cfg = &Config{}

type Config struct {
	ServerAddress string `yaml:"server_address" env-required:"true"`
	TgToken       string `yaml:"telegram_token" env-required:"true"`
	TgIdAllowed   int64  `yaml:"telegram_id_allowed" env-required:"true"`
	WgConfPath    string `yaml:"wireguard_config_path" env-default:"/etc/wireguard/wg0.conf"`
	WgServiceName string `yaml:"wireguard_service_name" env-default:"wg-quick@wg0.service"`
	WgPort        string `yaml:"wireguard_port" env-required:"true"`
	WgPubKey      string `yaml:"wireguard_publickey" env-required:"true"`
}

func GetConfig() *Config {

	once.Do(loadConfig)

	return Cfg
}

func loadConfig() {

	err := cleanenv.ReadConfig("config.yml", Cfg)
	if err != nil {
		log.Fatal(err)
	}

}
