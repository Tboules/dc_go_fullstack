package server

import (
	"fmt"
	"net/http"

	"github.com/Tboules/dc_go_fullstack/internal/auth"
	"github.com/Tboules/dc_go_fullstack/internal/constants"
	"github.com/Tboules/dc_go_fullstack/internal/views"
	"github.com/labstack/echo/v4"
)

func (s *Server) LoginPageHandler(c echo.Context) error {
	return views.LoginPage().Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) AuthProviderCallbackHandler(c echo.Context) error {
	user, err := s.auth.CompleteUserAuth(c)

	if err != nil {
		fmt.Println(err)
		return err
	}

	userClaims := auth.UserClaims{
		ProviderId: user.UserID,
	}

	accessToken, err := s.auth.NewAccessToken(userClaims)
	if err != nil {
		return err
	}

	refreshToken, err := s.auth.NewRefreshToken()
	if err != nil {
		return err
	}

	s.auth.AddTokenAsHttpOnlyCookie(accessToken, constants.AccessToken, c)
	s.auth.AddTokenAsHttpOnlyCookie(refreshToken, constants.RefreshToken, c)

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
