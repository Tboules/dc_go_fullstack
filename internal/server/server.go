package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Tboules/dc_go_fullstack/internal/auth"
	"github.com/Tboules/dc_go_fullstack/internal/database"
)

type Services struct {
	port  int
	store *database.Store
	auth  *auth.Auth
	DB    *database.SQLStore
}

func NewServer() (*http.Server, *Services) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	services := &Services{
		port:  port,
		store: database.New(),
		auth:  auth.InitAuth(),
		DB:    database.InitDatabase(),
	}

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      services.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return httpServer, services
}
