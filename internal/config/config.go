package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Cfg config

type config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	Server struct {
		Host string
		Port string
	}
}

func Load() error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	// Database Configs
	Cfg.Database.Host = getEnv("DATABASE_HOST", "localhost")
	Cfg.Database.Port = getEnv("DATABASE_PORT", "5432")
	Cfg.Database.User = getEnv("DATABASE_USER", "")
	Cfg.Database.Password = getEnv("DATABASE_PASSWORD", "")
	Cfg.Database.Name = getEnv("DATABASE_NAME", "")

	//Server Configs
	Cfg.Server.Host = getEnv("SERVER_HOST", "127.0.0.1:")
	Cfg.Server.Port = getEnv("SERVER_PORT", "8080")

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
