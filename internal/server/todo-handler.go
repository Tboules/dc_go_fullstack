package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Tboules/dc_go_fullstack/internal/database"
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

	deletedIndex := s.store.DeleteTodo(id)

	if deletedIndex == -1 {
		return c.String(http.StatusNotFound, "A todo at that id does not exist")
	}

	return c.NoContent(http.StatusOK)
}

func (s *Server) PostTodoHandler(c echo.Context) error {
	title := c.FormValue("title")
	desc := c.FormValue("description")

	todos := s.store.AddTodo(database.Todo{
		Title:       title,
		Description: desc,
	})

	comp := views.TodoCard(todos)

	return comp.Render(c.Request().Context(), c.Response().Writer)
}
