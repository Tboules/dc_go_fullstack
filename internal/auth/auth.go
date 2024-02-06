package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type Auth struct {
	Store sessions.Store
}

const (
	MaxAge = 86400 * 30
	IsProd = false
)

func InitAuth() *Auth {

	key := []byte(os.Getenv("AUTH_SESSION_SECRET"))

	store := sessions.NewCookieStore(key)
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8080/auth/google/callback"),
	)

	return &Auth{
		Store: store,
	}
}
