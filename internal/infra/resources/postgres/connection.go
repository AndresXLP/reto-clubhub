package postgres

import (
	"fmt"
	"sync"

	"franchises-system/config"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

var (
	connection *sqlx.DB
	once       sync.Once
)

func NewPostgresConnection() *sqlx.DB {
	once.Do(func() {
		connection = getConnection()
	})
	return connection
}

func getConnection() *sqlx.DB {
	urlConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Environments().Postgres.Host,
		config.Environments().Postgres.Port,
		config.Environments().Postgres.User,
		config.Environments().Postgres.Password,
		config.Environments().Postgres.DbName)

	db, err := sqlx.Connect("postgres", urlConnection)
	if err != nil {
		log.Panic(err)
	}

	log.Info("Postgres Successfully Connected")
	return db
}
