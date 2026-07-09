package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Cfg config

type config struct {
	Database struct {
		Host            string
		Port            string
		User            string
		Password        string
		Name            string
		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxLifetime int
		ConnMaxIdleTime int
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
	Cfg.Database.MaxOpenConns = getEnvAsInt("MAX_OPEN_CONNS", 20)
	Cfg.Database.MaxIdleConns = getEnvAsInt("MAX_IDLE_CONNS", 5)
	Cfg.Database.ConnMaxLifetime = getEnvAsInt("CONN_MAX_LIFETIME", 30)
	Cfg.Database.ConnMaxIdleTime = getEnvAsInt("CONN_MAX_IDLE_TIME", 2)

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

func getEnvAsInt(key string, fallback int) int {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Invalid value for %s. Using default %d", key, fallback)
		return fallback
	}

	return value
}
