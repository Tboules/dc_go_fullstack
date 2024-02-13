package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/Tboules/dc_go_fullstack/internal/database/sqlc"
	_ "github.com/go-sql-driver/mysql"
)

type SQLStore struct {
	Conn    *sql.DB
	Queries *sqlc.Queries
}

func InitDatabase() *SQLStore {
	conString := os.Getenv("DSN")

	db, err := sql.Open("mysql", conString)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Problem pinging the db: %v", err)
	}

	log.Println("Db Connection Opened ")

	store := &SQLStore{
		Conn:    db,
		Queries: sqlc.New(db),
	}

	return store
}
