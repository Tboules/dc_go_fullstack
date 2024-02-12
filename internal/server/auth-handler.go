package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Tboules/dc_go_fullstack/internal/auth"
	"github.com/Tboules/dc_go_fullstack/internal/constants"
	"github.com/Tboules/dc_go_fullstack/internal/database/sqlc"
	"github.com/Tboules/dc_go_fullstack/internal/views"
	"github.com/labstack/echo/v4"
)

func (s *Server) LoginPageHandler(c echo.Context) error {
	return views.LoginPage().Render(c.Request().Context(), c.Response().Writer)
}

func (s *Server) AuthProviderCallbackHandler(c echo.Context) error {
	user, err := s.auth.CompleteUserAuth(c)
	if err != nil {
		fmt.Println(err)
		return err
	}

	userClaims := auth.UserClaims{
		Email:      user.Email,
		ProviderId: user.UserID,
	}

	appUser, err := s.db.Queries.GetUserByProviderId(c.Request().Context(), user.UserID)
	if err != nil {
		fmt.Println("creating new user")
		fmt.Println(user.AvatarURL)
		fmt.Println(user.Name)
		fmt.Println(user.FirstName)
		fmt.Println(user)

		//name is null
		//fix

		res, err := s.db.Queries.CreateNewUser(c.Request().Context(), sqlc.CreateNewUserParams{
			Email: user.Email,
			Name:  user.FirstName + user.LastName,
			Image: sql.NullString{String: user.AvatarURL, Valid: true},
		})
		if err != nil {
			log.Println("Error adding user")
		}

		userId, err := res.LastInsertId()
		if err != nil {
			log.Println("Error getting last id")
		}

		userClaims.UserID = userId
	} else {
		userClaims.UserID = appUser.ID
	}

	accessToken, err := s.auth.NewAccessToken(userClaims)
	if err != nil {
		return err
	}

	refreshToken, err := s.auth.NewRefreshToken()
	if err != nil {
		return err
	}

	//create session with refresh token
	_, err = s.db.Queries.SaveSession(c.Request().Context(), sqlc.SaveSessionParams{
		Token:     refreshToken,
		UserID:    userClaims.UserID,
		ExpiresAt: s.auth.NewRefreshExpiry(),
	})
	if err != nil {
		log.Printf("Error creating session: %v", err)
	}

	s.auth.AddTokenAsHttpOnlyCookie(accessToken, constants.AccessToken, c)
	s.auth.AddTokenAsHttpOnlyCookie(refreshToken, constants.RefreshToken, c)

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
