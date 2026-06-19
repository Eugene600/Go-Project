package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Eugene600/Go-Project/internal/config"
	"github.com/Eugene600/Go-Project/internal/database"
	"github.com/Eugene600/Go-Project/internal/routes"
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

	defer db.Close()

	log.Printf("Database connection successful")

	handler := routes.SetRoutes()

	server := &http.Server{
		Addr:    config.Cfg.Server.Host + config.Cfg.Server.Port,
		Handler: handler,
	}

	log.Println("Server running on: ", config.Cfg.Server.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
