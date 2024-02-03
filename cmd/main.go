package main

import (
	"log"
	"os"

	"github.com/Tboules/dc_go_fullstack/internal/server"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8080/auth/google/callback"),
	)

	server := server.NewServer()

	log.Fatal(server.ListenAndServe())
}
