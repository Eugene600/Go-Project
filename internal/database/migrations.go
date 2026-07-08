package database

import (
	"log"

	"github.com/pressly/goose/v3"
)

type MigrationDirection string

const (
	MigrateUp     MigrationDirection = "up"
	MigrateDown   MigrationDirection = "down"
	MigrateStatus MigrationDirection = "status"
)

func Migrate(direction MigrationDirection) {
	err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()

	switch direction {
	case MigrateUp:
		err = goose.Up(DB, "internal/migrations")
	case MigrateDown:
		err = goose.Down(DB, "internal/migrations")
	case MigrateStatus:
		err = goose.Status(DB, "internal/migrations")
	default:
		log.Fatal("Unknown migration direction")
	}

	if err != nil {
		log.Fatal(err)
	}
}
