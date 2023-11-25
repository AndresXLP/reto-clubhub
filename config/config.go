package config

import (
	"sync"

	"github.com/andresxlp/gosuite/config"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Server   Server   `mapstructure:"server" validate:"required"`
	Postgres Postgres `mapstructure:"postgres" validate:"required"`
}

type Server struct {
	Port int `mapstructure:"port" validate:"required"`
}

type Postgres struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	DbName   string `mapstructure:"db_name" validate:"required"`
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
