package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Tboules/dc_go_fullstack/internal/auth"
	"github.com/Tboules/dc_go_fullstack/internal/constants"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (s *Server) AuthProviderCallbackHandler(c echo.Context) error {
	user, err := s.auth.CompleteUserAuth(c)

	if err != nil {
		fmt.Println(err)
		return err
	}

	userClaims := auth.UserClaims{
		ProviderId: user.UserID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	accessToken, err := s.auth.NewAccessToken(userClaims)
	if err != nil {
		return err
	}

	refreshToken, err := s.auth.NewRefreshToken()
	if err != nil {
		return err
	}

	accessCookie := new(http.Cookie)
	accessCookie.Name = constants.AccessToken
	accessCookie.Value = accessToken
	accessCookie.HttpOnly = true
	accessCookie.Secure = false
	accessCookie.Path = "/"

	c.SetCookie(accessCookie)

	refreshCookie := new(http.Cookie)
	refreshCookie.Name = constants.RefreshToken
	refreshCookie.Value = refreshToken
	refreshCookie.HttpOnly = true
	refreshCookie.Secure = false
	refreshCookie.Path = "/"

	c.SetCookie(refreshCookie)

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
