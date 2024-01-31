package server

import (
	"github.com/Tboules/dc_go_fullstack/internal/views"
	"github.com/labstack/echo/v4"
)

func (s *Server) HomeHandler(c echo.Context) error {
	comp := views.HomePage(s.store.CurrentCount())

	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) PostCount(c echo.Context) error {
	count := s.store.Increment()

	comp := views.CountButton(count)

	return comp.Render(c.Request().Context(), c.Response().Writer)
}
