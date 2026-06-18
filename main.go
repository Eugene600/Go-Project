package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/Eugene600/Go-Project/internal/config"
	"github.com/Eugene600/Go-Project/internal/database"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Error Loading Configs %s", err)
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

	db, err := database.Connect(connString)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Database connection successful")

	defer db.Close()
}
