package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Services) RegisterRoutes() *echo.Echo {
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

func (s *Services) homeRouter(e *echo.Echo) {
	e.GET("/", s.HomeHandler, s.unrestrictedSetUserClaims)
	e.POST("/", s.PostCount)
}

func (s *Services) authRouter(e *echo.Echo) {
	e.GET("/login", s.LoginPageHandler)
	e.GET("/auth/:provider/callback", s.AuthProviderCallbackHandler)
	e.GET("/auth/:provider", s.AuthHandler)

	e.GET("/auth/logout", s.LogoutHandler)
}

// secure routes
func (s *Services) todoRouter(e *echo.Echo) {
	todoGroup := e.Group("/todo", s.secureRoutesMiddleware)

	todoGroup.GET("", s.TodoPageHandler)
	todoGroup.POST("", s.PostTodoHandler)

	todoGroup.DELETE("/:id", s.DeleteTodoHandler)
}
