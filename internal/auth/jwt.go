package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/Tboules/dc_go_fullstack/internal/utils"
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	ProviderId string `json:"provider_id"`
	UserID     int64  `json:"id"`
	Email      string `json:"email"`
	jwt.StandardClaims
}

func (a *Auth) NewAccessToken(claims UserClaims) (string, error) {
	claims.StandardClaims = jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: utils.NewAccessExpiry().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SIGNING_SECRET")))
}

func (a *Auth) NewRefreshToken() (string, error) {
	claims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: utils.NewRefreshExpiry().Unix(),
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
