package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/Eugene600/Go-Project/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {
	if err := config.Load(); err != nil {
		return nil, err
	}

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.Cfg.Database.User,
		url.QueryEscape(config.Cfg.Database.Password),
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		config.Cfg.Database.Name,
	)

	log.Println("Connection string is ", connString)

	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected Successfully to DB")

	return db, nil
}
