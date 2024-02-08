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
	ProviderId string `json:"id"`
	jwt.StandardClaims
}

func (a *Auth) NewAccessToken(claims UserClaims) (string, error) {
	claims.StandardClaims = jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SIGNING_SECRET")))
}

func (a *Auth) NewRefreshToken() (string, error) {
	claims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SIGNING_SECRET")))
}

func (a *Auth) ParseAccessToken(token string) (*UserClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return &UserClaims{}, fmt.Errorf("there was something wrong with the token")
		}

		return []byte(os.Getenv("JWT_SIGNING_SECRET")), nil
	})

	if err != nil {
		fmt.Println(err)
		return &UserClaims{}, err
	}

	if claims, ok := parsedAccessToken.Claims.(*UserClaims); ok && parsedAccessToken.Valid {
		return claims, nil
	} else {
		return &UserClaims{}, fmt.Errorf("Invalid Token")
	}
}

func (a *Auth) ParseRefreshToken(token string) (*jwt.StandardClaims, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return &jwt.StandardClaims{}, nil
		}

		return []byte(os.Getenv("JWT_SIGNING_SECRET")), nil
	})

	if err != nil {
		fmt.Println(err)
		return &jwt.StandardClaims{}, err
	}

	claims, ok := parsedRefreshToken.Claims.(*jwt.StandardClaims)

	if ok && parsedRefreshToken.Valid {
		return claims, nil
	} else {
		return &jwt.StandardClaims{}, fmt.Errorf("Invalid Refresh Token")
	}
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
