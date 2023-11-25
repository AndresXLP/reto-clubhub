package config

import (
	"sync"

	"github.com/andresxlp/gosuite/config"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Server Server `mapstructure:"server" validate:"required"`
}

type Server struct {
	Port int `mapstructure:"port" validate:"required"`
}

var (
	once sync.Once
	Cfg  Config
)

func Environments() Config {
	once.Do(func() {
		if err := config.GetConfigFromEnv(&Cfg); err != nil {
			log.Panic(err)
		}
	})

	return Cfg
}
