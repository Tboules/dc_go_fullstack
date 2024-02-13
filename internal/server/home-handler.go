package server

import (
	"fmt"

	"github.com/Tboules/dc_go_fullstack/internal/auth"
	"github.com/Tboules/dc_go_fullstack/internal/constants"
	"github.com/Tboules/dc_go_fullstack/internal/views"
	"github.com/labstack/echo/v4"
)

func (s *Services) HomeHandler(c echo.Context) error {
	claims, _ := c.Get(constants.UserClaimsKey).(*auth.UserClaims)
	fmt.Printf("claims: %+v \n", claims)

	comp := views.HomePage(s.store.CurrentCount())

	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func (s *Services) PostCount(c echo.Context) error {
	count := s.store.Increment()

	comp := views.CountButton(count)

	return comp.Render(c.Request().Context(), c.Response().Writer)
}
