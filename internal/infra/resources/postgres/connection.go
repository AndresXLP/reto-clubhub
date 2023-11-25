package postgres

import (
	"fmt"
	"sync"

	"franchises-system/config"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	connection *gorm.DB
	once       sync.Once
)

func NewPostgresConnection() *gorm.DB {
	once.Do(func() {
		connection = getConnection()
	})
	return connection
}

func getConnection() *gorm.DB {
	urlConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Environments().Postgres.Host,
		config.Environments().Postgres.Port,
		config.Environments().Postgres.User,
		config.Environments().Postgres.Password,
		config.Environments().Postgres.DbName)

	db, err := gorm.Open(postgres.Open(urlConnection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panic(err)
	}

	log.Info("Postgres Successfully Connected")
	return db
}
