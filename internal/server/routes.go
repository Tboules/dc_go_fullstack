package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	e.GET("/", homeHandler)

	return e
}

func homeHandler(ctx echo.Context) error {
	resp := make(map[string]string)

	resp["Message"] = "Hello Air"

	return ctx.JSON(http.StatusOK, resp)
}
