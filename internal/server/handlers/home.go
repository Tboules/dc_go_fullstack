package handlers

import (
	"net/http"
	"strconv"

	"github.com/Tboules/dc_go_fullstack/internal/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Demo struct {
	count int
}

var d = &Demo{
	count: 0,
}

func HomeHandler() echo.HandlerFunc {

	comp := views.HomePage("Templ HTML Template")

	return echo.WrapHandler(templ.Handler(comp))
}

func PostCount(c echo.Context) error {
	d.count = d.count + 1

	return c.String(http.StatusOK, strconv.Itoa(d.count))
}
