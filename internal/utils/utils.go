package utils

import (
	"net/http"
	"time"

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

func NewAccessExpiry() time.Time {
	return time.Now().Add(time.Minute * 15)
}

func NewRefreshExpiry() time.Time {
	return time.Now().Add(time.Hour * 48)
}
