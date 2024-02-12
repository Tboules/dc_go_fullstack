package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide correct args for migration")
	}

	arg := os.Args[1]

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	driver, _ := mysql.WithInstance(db, &mysql.Config{})

	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"mysql",
		driver,
	)

	if arg == "up" {
		log.Println(arg)
		err = m.Up()
		if err != nil {
			log.Fatalf("Problem with up migration: %v", err)
		}
	} else if arg == "down" {
		log.Println(arg)
		err = m.Down()
		if err != nil {
			log.Fatalf("Problem with down migration: %v", err)
		}
	} else {
		log.Fatal("Must provide up or down command")
	}
}
