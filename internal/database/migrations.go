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
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	switch direction {
	case MigrateUp:
		err = goose.Up(db, "internal/migrations")
	case MigrateDown:
		err = goose.Down(db, "internal/migrations")
	case MigrateStatus:
		err = goose.Status(db, "internal/migrations")
	default:
		log.Fatal("Unknown migration direction")
	}

	if err != nil {
		log.Fatal(err)
	}
}
