package server

import (
	"fmt"
	"net/http"

	"github.com/Tboules/dc_go_fullstack/internal/constants"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = customErrorHandler

	e.Static("/static", "cmd/web")
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	s.authRouter(e)
	s.homeRouter(e)

	s.todoRouter(e)

	return e
}

// open routes

func (s *Server) homeRouter(e *echo.Echo) {
	e.GET("/", s.HomeHandler)
	e.POST("/", s.PostCount)
}

func (s *Server) authRouter(e *echo.Echo) {
	e.GET("/login", s.LoginPageHandler)
	e.GET("/auth/:provider/callback", s.AuthProviderCallbackHandler)
	e.GET("/auth/:provider", s.AuthHandler)

	e.GET("/auth/logout", s.LogoutHandler)
}

// secure routes

func (s *Server) todoRouter(e *echo.Echo) {
	todoGroup := e.Group("/todo", s.secureRoutesMiddleware)

	todoGroup.GET("", s.TodoPageHandler)
	todoGroup.POST("", s.PostTodoHandler)

	todoGroup.DELETE("/:id", s.DeleteTodoHandler)
}

// middleware
func (s *Server) secureRoutesMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken, err := c.Cookie(constants.AccessToken)
		if err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusUnauthorized, "No access token found in cookies")
		}

		claims, err := s.auth.ParseAccessToken(accessToken.Value)

		if err != nil || claims.Valid() != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid access token")
		}

		return next(c)
	}
}

// global error handler
func customErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)

	switch code {
	case http.StatusUnauthorized:
		err := c.Redirect(http.StatusTemporaryRedirect, "/login")
		if err != nil {
			fmt.Println("Problem with redirect in global error handler")
		}
	}
}
