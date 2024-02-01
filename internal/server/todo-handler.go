package server

import (
	"github.com/Tboules/dc_go_fullstack/internal/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func (s *Server) TodoPageHandler() echo.HandlerFunc {
	page := views.TodoPage(s.store.GetTodos())

	return echo.WrapHandler(templ.Handler(page))
}
