package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	ProviderId string `json:"id"`
	jwt.StandardClaims
}

func (a *Auth) NewAccessToken(claims UserClaims) (string, error) {
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
		return []byte(os.Getenv("JWT_SIGNING_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return parsedAccessToken.Claims.(*UserClaims), nil
}

func (a *Auth) ParseRefreshToken(token string) (*jwt.StandardClaims, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SIGNING_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedRefreshToken.Claims.(*jwt.StandardClaims), nil
}
