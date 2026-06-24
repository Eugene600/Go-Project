package main

import (
	"log"
	"net/http"

	"github.com/Eugene600/Go-Project/internal/config"
	"github.com/Eugene600/Go-Project/internal/routes"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	
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
