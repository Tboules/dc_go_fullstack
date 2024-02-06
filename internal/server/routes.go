package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() *echo.Echo {
	e := echo.New()

	e.Static("/static", "cmd/web")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	s.authRouter(e)
	s.homeRouter(e)
	s.todoRouter(e)

	return e
}

func (s *Server) homeRouter(e *echo.Echo) {
	e.GET("/", s.HomeHandler)
	e.POST("/", s.PostCount)
}

func (s *Server) todoRouter(e *echo.Echo) {
	e.GET("/todo", s.TodoPageHandler)
	e.DELETE("/todo/:id", s.DeleteTodoHandler)
	e.POST("/todo", s.PostTodoHandler)
}

func (s *Server) authRouter(e *echo.Echo) {
	e.GET("/auth/:provider/callback", s.AuthProviderCallbackHandler)
	e.GET("/auth/:provider", s.AuthHandler)

	e.GET("/auth/logout", s.LogoutHandler)
}
