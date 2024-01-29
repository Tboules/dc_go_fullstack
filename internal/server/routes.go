package server

import (
	"github.com/Tboules/dc_go_fullstack/internal/views"
	"github.com/a-h/templ"
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

	e.GET("/", homeHandler())

	return e
}

func homeHandler() echo.HandlerFunc {
	comp := views.HomePage("Templ HTML Template")

	return echo.WrapHandler(templ.Handler(comp))
}
