package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserClaims struct {
	ProviderId string `json:"provider_id"`
	UserID     int64  `json:"id"`
	Email      string `json:"email"`
	jwt.StandardClaims
}

func (a *Auth) NewAccessExpiry() time.Time {
	return time.Now().Add(time.Minute * 15)
}

func (a *Auth) NewRefreshExpiry() time.Time {
	return time.Now().Add(time.Hour * 48)
}

func (a *Auth) NewAccessToken(claims UserClaims) (string, error) {
	claims.StandardClaims = jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: a.NewAccessExpiry().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SIGNING_SECRET")))
}

func (a *Auth) NewRefreshToken() (string, error) {
	claims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: a.NewRefreshExpiry().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SIGNING_SECRET")))
}

func (a *Auth) ParseAccessToken(token string) (*UserClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNING_SECRET")), nil
	})

	if err != nil {
		if err.(*jwt.ValidationError).Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("Token Malformed")
			return &UserClaims{}, err
		}
	}

	return parsedAccessToken.Claims.(*UserClaims), nil
}

func (a *Auth) ParseRefreshToken(token string) (*jwt.StandardClaims, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNING_SECRET")), nil
	})

	if err != nil {
		if err.(*jwt.ValidationError).Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("Refresh Token Malformed")
			return &jwt.StandardClaims{}, err
		}
	}

	return parsedRefreshToken.Claims.(*jwt.StandardClaims), nil
}

func (a *Auth) AddTokenAsHttpOnlyCookie(token string, key string, c echo.Context) {
	accessCookie := new(http.Cookie)
	accessCookie.Name = key
	accessCookie.Value = token
	accessCookie.HttpOnly = true
	accessCookie.Secure = false
	accessCookie.Path = "/"

	c.SetCookie(accessCookie)
}
