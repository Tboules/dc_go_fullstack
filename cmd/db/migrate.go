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

	err = m.Up()
	if err != nil {
		log.Fatalf("Problem with up migration: %v", err)
	}
	log.Println("Migration completed sucessfully")
}
