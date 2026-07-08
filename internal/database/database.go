package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/Eugene600/Go-Project/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() error {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.Cfg.Database.User,
		url.QueryEscape(config.Cfg.Database.Password),
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		config.Cfg.Database.Name,
	)

	log.Println("Connection string is ", connString)

	var err error

	DB, err = sql.Open("pgx", connString)
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}

	log.Println("Connected Successfully to DB")

	return nil
}
