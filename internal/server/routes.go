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
	e.GET("/", s.HomeHandler, s.unrestrictedSetUserClaims)
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

		// get tokens from cookies
		accessToken, err := c.Cookie(constants.AccessToken)
		if err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusUnauthorized, "No access token found in cookies")
		}
		refreshToken, err := c.Cookie(constants.RefreshToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "No refresh token found")
		}

		// parse tokens
		userClaims, ucError := s.auth.ParseAccessToken(accessToken.Value)
		refreshClaims, rcError := s.auth.ParseRefreshToken(refreshToken.Value)

		// check if token format is messed up
		if ucError != nil || rcError != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "error with token format")
		}

		// refresh refresh token if expired
		if refreshClaims.Valid() != nil {
			// check that old token exists in sessions table
			// if not redirect user to sign in again
			//delete old token from db token and/or sessions table

			freshRefresheToken, err := s.auth.NewRefreshToken()
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "error creating refresh token")
			}

			s.auth.AddTokenAsHttpOnlyCookie(freshRefresheToken, constants.RefreshToken, c)
		}

		// refresh access token if refresh token is valid
		if userClaims.StandardClaims.Valid() != nil && refreshClaims.Valid() == nil {
			freshAccessToken, err := s.auth.NewAccessToken(*userClaims)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "error creating access token")
			}
			s.auth.AddTokenAsHttpOnlyCookie(freshAccessToken, constants.AccessToken, c)

			fmt.Printf("old expiration: %v\n", userClaims.StandardClaims.ExpiresAt)
			fmt.Printf("new token: %s\n", freshAccessToken)
		}

		c.Set(constants.UserClaimsKey, userClaims)
		return next(c)
	}
}

func (s *Server) unrestrictedSetUserClaims(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken, _ := c.Cookie(constants.AccessToken)

		claimsFromToken, _ := s.auth.ParseAccessToken(accessToken.Value)

		claims := claimsFromToken

		c.Set(constants.UserClaimsKey, claims)
		return next(c)
	}
}

// global error handler
func customErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	switch code {
	case http.StatusUnauthorized:
		err := c.Redirect(http.StatusTemporaryRedirect, "/login")
		if err != nil {
			fmt.Println("Problem with redirect in global error handler")
		}
	}
}
