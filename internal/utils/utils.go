package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Tboules/dc_go_fullstack/internal/constants"
	"github.com/labstack/echo/v4"
)

func AddHttpOnlyCookie(key string, value string, c echo.Context) {
	accessCookie := new(http.Cookie)
	accessCookie.Name = key
	accessCookie.Value = value
	accessCookie.HttpOnly = true
	accessCookie.Secure = false
	accessCookie.Path = "/"

	c.SetCookie(accessCookie)
}

func ClearAuthCookies(c echo.Context) error {
	accessCookie, err := c.Cookie(constants.AccessToken)
	if err != nil {
		return err
	}
	refreshCookie, err := c.Cookie(constants.RefreshToken)
	if err != nil {
		return err
	}

	accessCookie.Path = "/"
	accessCookie.HttpOnly = true
	accessCookie.MaxAge = -1
	refreshCookie.Path = "/"
	refreshCookie.HttpOnly = true
	refreshCookie.MaxAge = -1

	c.SetCookie(accessCookie)
	c.SetCookie(refreshCookie)

	return nil
}

func NewAccessExpiry() time.Time {
	return time.Now().Add(time.Minute * 15)
}

func NewRefreshExpiry() time.Time {
	return time.Now().Add(time.Hour * 48)
}
