package main

import (
	"log"

	"github.com/Tboules/dc_go_fullstack/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := server.NewServer()

	log.Fatal(server.ListenAndServe())
}
