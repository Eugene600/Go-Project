package main

import (
    "fmt"
    "os"

    "github.com/Eugene600/Go-Project/internal/database"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run ./cmd/migrate [up|down|status]")
        return
    }

    switch os.Args[1] {
    case "up":
        database.Migrate(database.MigrateUp)

    case "down":
        database.Migrate(database.MigrateDown)

    case "status":
        database.Migrate(database.MigrateStatus)

    default:
        fmt.Println("Unknown command")
    }
}