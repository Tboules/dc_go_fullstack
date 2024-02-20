package server

import (
	"fmt"
	"net/http"

	"github.com/Tboules/dc_go_fullstack/internal/constants"
	"github.com/Tboules/dc_go_fullstack/internal/database/sqlc"
	"github.com/Tboules/dc_go_fullstack/internal/utils"
	"github.com/labstack/echo/v4"
)

func (s *Services) secureRoutesMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
			// check that old token exists
			_, err := s.DB.Queries.GetSession(c.Request().Context(), refreshToken.Value)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token not found")
			}

			// delete old token from db token and/or sessions table
			err = s.DB.Queries.DeleteSession(c.Request().Context(), refreshToken.Value)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token not deleted")
			}

			freshRefresheToken, err := s.auth.NewRefreshToken()
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "error creating refresh token")
			}

			_, err = s.DB.Queries.SaveSession(c.Request().Context(), sqlc.SaveSessionParams{
				Token:     freshRefresheToken,
				UserID:    userClaims.UserID,
				ExpiresAt: utils.NewRefreshExpiry(),
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "error creating session")
			}

			// add token in sessions table
			utils.AddHttpOnlyCookie(constants.RefreshToken, freshRefresheToken, c)
		}

		// refresh access token if refresh token is valid
		if userClaims.StandardClaims.Valid() != nil && refreshClaims.Valid() == nil {
			freshAccessToken, err := s.auth.NewAccessToken(*userClaims)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "error creating access token")
			}
			utils.AddHttpOnlyCookie(constants.AccessToken, freshAccessToken, c)

			fmt.Printf("old expiration: %v\n", userClaims.StandardClaims.ExpiresAt)
			fmt.Printf("new token: %s\n", freshAccessToken)
		}

		c.Set(constants.UserClaimsKey, userClaims)
		return next(c)
	}
}

func (s *Services) unrestrictedSetUserClaims(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken, err := c.Cookie(constants.AccessToken)
		if err != nil {
			return next(c)
		}

		claimsFromToken, err := s.auth.ParseAccessToken(accessToken.Value)
		if err != nil {
			return next(c)
		}

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
