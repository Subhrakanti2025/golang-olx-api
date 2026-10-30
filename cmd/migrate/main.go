package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Subhrakanti2025/golang-olx-api/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	fmt.Println("runnign migration", os.Args)

	// Check whether the user provided a migration command.
	// Example: up, down, or force.
	if len(os.Args) < 2 {
		log.Fatal("Mention migration type like <up | down | force>")
	}

	cfg := config.MustLoad()

	// Initialize the migration instance.
	// "file://migrations" specifies the directory containing .sql migration files.
	// cfg.DatabaseUrl provides the database connection URL.
	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()
	switch os.Args[1] {

	// CASE 1: UP MIGRATION
	// Execute all pending UP migrations.
	// It applies database changes that have not yet been executed.
	// Example: go run ./cmd/migrate/main.go up
	case "up":
		fmt.Println("migration up")
		if err := m.Up(); err != nil {
			log.Fatal("migration up error", err)
		}

	// CASE 2: DOWN MIGRATION
	// Roll back the most recently applied migration.
	// m.Steps(-1) means execute one DOWN migration.
	// Example: go run ./cmd/migrate/main.go down
	case "down":
		fmt.Println("migration down")
		if err := m.Steps(-1); err != nil {
			log.Fatal("migration down error", err)
		}

	// CASE 3: FORCE MIGRATION
	// Manually change the migration version stored in the database.
	// Useful when resolving a dirty migration state after inspecting and
	// reconciling the actual database schema.
	// Example: go run ./cmd/migrate/main.go force 1
	// os.Args[2] contains the version as a string ("1").
	// strconv.Atoi() converts that string into an integer (1).
	// m.Force() updates migration tracking; it DOES NOT execute SQL files
	// or repair the actual database schema.
	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Mention migration version, example: force <version>")
		}
		fmt.Println("migration force", os.Args[2])
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("Invalid migration version")
		}
		if err := m.Force(version); err != nil {
			log.Fatal("migration force error: ", err)
		}

	// DEFAULT CASE
	// Executes when the user provides an unsupported migration command.
	// Example: go run ./cmd/migrate/main.go something
	default:
		log.Fatal("Mention your migration type")
	}
}
