package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Tboules/dc_go_fullstack/internal/auth"
	"github.com/Tboules/dc_go_fullstack/internal/constants"
	"github.com/Tboules/dc_go_fullstack/internal/database"
	"github.com/Tboules/dc_go_fullstack/internal/views"
	"github.com/labstack/echo/v4"
)

func (s *Services) TodoPageHandler(c echo.Context) error {
	claims, _ := c.Get(constants.UserClaimsKey).(*auth.UserClaims)
	page := views.TodoPage(s.store.GetTodos(), claims)

	return page.Render(c.Request().Context(), c.Response().Writer)
}

func (s *Services) DeleteTodoHandler(c echo.Context) error {
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

func (s *Services) PostTodoHandler(c echo.Context) error {
	title := c.FormValue("title")
	desc := c.FormValue("description")

	todos := s.store.AddTodo(database.Todo{
		Title:       title,
		Description: desc,
	})

	comp := views.TodoCard(todos)

	return comp.Render(c.Request().Context(), c.Response().Writer)
}
