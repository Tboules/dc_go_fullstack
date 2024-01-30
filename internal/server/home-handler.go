package server

import (
	"github.com/Tboules/dc_go_fullstack/internal/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func (s *Server) HomeHandler() echo.HandlerFunc {
	comp := views.HomePage("Templ HTML Template")

	return echo.WrapHandler(templ.Handler(comp))
}

func (s *Server) PostCount(c echo.Context) error {
	comp := views.CountButton(s.store.Increment())

	return comp.Render(c.Request().Context(), c.Response().Writer)
}
