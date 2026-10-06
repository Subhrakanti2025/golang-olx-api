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
	if len(os.Args) < 2 {
		log.Fatal("Mention migration type like <up | down | force>")
	}

	cfg := config.MustLoad()
	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()
	switch os.Args[1] {
	case "up":
		fmt.Println("migration up")
		if err := m.Up(); err != nil {
			log.Fatal("migration up error", err)
		}
	case "down":
		fmt.Println("migration down")
		if err := m.Down(); err != nil {
			log.Fatal("migration down error", err)
		}

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

	default:
		log.Fatal("Mention your migration type")
	}
}
