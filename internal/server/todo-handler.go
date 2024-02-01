package server

import (
	"errors"
	"strconv"

	"github.com/Tboules/dc_go_fullstack/internal/views"
	"github.com/labstack/echo/v4"
)

func (s *Server) TodoPageHandler(c echo.Context) error {
	page := views.TodoPage(s.store.GetTodos())

	return page.Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) DeleteTodoHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return errors.New("No ID was provided")
	}

	todos := s.store.DeleteTodo(id)

	comp := views.TodoCard(todos)

	return comp.Render(c.Request().Context(), c.Response().Writer)
}
