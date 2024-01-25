package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string
}

func AppRouter() *echo.Echo {
	router := echo.New()

	router.GET("/", homeHandler)

	return router
}

func homeHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK,
		Response{Message: "Hello World"},
	)
}
