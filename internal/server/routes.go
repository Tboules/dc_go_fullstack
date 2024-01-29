package server

import (
	"github.com/Tboules/dc_go_fullstack/internal/server/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() *echo.Echo {
	e := echo.New()

	e.Static("/static", "cmd/web")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	e.GET("/", handlers.HomeHandler())
	e.POST("/", handlers.PostCount)

	return e
}
