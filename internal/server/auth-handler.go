package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Tboules/dc_go_fullstack/internal/auth"
	"github.com/Tboules/dc_go_fullstack/internal/constants"
	"github.com/Tboules/dc_go_fullstack/internal/database/sqlc"
	"github.com/Tboules/dc_go_fullstack/internal/utils"
	"github.com/Tboules/dc_go_fullstack/internal/views"
	"github.com/labstack/echo/v4"
)

func (s *Services) LoginPageHandler(c echo.Context) error {
	return views.LoginPage().Render(c.Request().Context(), c.Response().Writer)
}

func (s *Services) AuthProviderCallbackHandler(c echo.Context) error {
	user, err := s.auth.CompleteUserAuth(c)
	if err != nil {
		fmt.Println(err)
		return err
	}

	userClaims := auth.UserClaims{
		Email:      user.Email,
		ProviderId: user.UserID,
	}

	appUser, err := s.DB.Queries.GetUserByProviderId(c.Request().Context(), user.UserID)
	if err != nil {
		name := user.FirstName + user.LastName

		res, err := s.DB.Queries.CreateNewUser(c.Request().Context(), sqlc.CreateNewUserParams{
			Email:      user.Email,
			Name:       sql.NullString{String: name, Valid: name != ""},
			Image:      sql.NullString{String: user.AvatarURL, Valid: true},
			ProviderID: user.UserID,
		})
		if err != nil {
			log.Printf("Error adding user %v\n", err)
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
	_, err = s.DB.Queries.SaveSession(c.Request().Context(), sqlc.SaveSessionParams{
		Token:     refreshToken,
		UserID:    userClaims.UserID,
		ExpiresAt: utils.NewRefreshExpiry(),
	})
	if err != nil {
		log.Printf("Error creating session: %v", err)
	}

	utils.AddHttpOnlyCookie(constants.AccessToken, accessToken, c)
	utils.AddHttpOnlyCookie(constants.RefreshToken, refreshToken, c)

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (s *Services) AuthHandler(c echo.Context) error {
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

func (s *Services) LogoutHandler(c echo.Context) error {
	err := s.auth.Logout(c)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
