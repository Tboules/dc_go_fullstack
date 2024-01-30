package server

import (
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

	homeRouter(e, s)

	return e
}

func homeRouter(e *echo.Echo, s *Server) {
	e.GET("/", s.HomeHandler())
	e.POST("/", s.PostCount)
}
