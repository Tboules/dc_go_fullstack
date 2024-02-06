package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) AuthProviderCallbackHandler(c echo.Context) error {
	user, err := s.auth.CompleteUserAuth(c)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(user)

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (s *Server) AuthHandler(c echo.Context) error {
	gothUser, err := s.auth.CompleteUserAuth(c)

	if err == nil {
		fmt.Printf("user already exists: %v", gothUser)
		return nil
	} else {
		err := s.auth.BeginAuthHandler(c)
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

func (s *Server) LogoutHandler(c echo.Context) error {
	err := s.auth.Logout(c)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
