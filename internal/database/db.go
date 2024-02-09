package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	conString := os.Getenv("DSN")

	db, err := sql.Open("mysql", conString)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatalf("Failed to query: %v", err)
	}
	defer rows.Close()

	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}

		log.Println(tableName)
	}

	defer db.Close()

	return db
}
