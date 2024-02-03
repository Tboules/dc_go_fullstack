package auth

// copy of echo-gothic repo code
// github.com/nabowler/echo-gothic

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func (a *Auth) BeginAuthHandler(c echo.Context) error {
	url, err := a.GetAuthURL(c)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, url)

}

func (a *Auth) GetAuthURL(c echo.Context) (string, error) {
	providerName, err := a.GetProviderName(c)
	if err != nil {
		return "", err
	}

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return "", err
	}

	sess, err := provider.BeginAuth(a.SetState(c))
	if err != nil {
		return "", err
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}

	err = a.StoreInSession(providerName, sess.Marshal(), c)
	if err != nil {
		return "", err
	}

	return url, err
}

func (a *Auth) CompleteUserAuth(c echo.Context) (goth.User, error) {
	defer func() {
		_ = a.Logout(c)
	}()

	providerName, err := a.GetProviderName(c)
	if err != nil {
		return goth.User{}, err
	}

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return goth.User{}, err
	}

	value, err := a.GetFromSession(providerName, c)
	if err != nil {
		return goth.User{}, err
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return goth.User{}, err
	}

	err = a.validateState(c, sess)
	if err != nil {
		return goth.User{}, err
	}

	user, err := provider.FetchUser(sess)
	if err == nil {
		return user, err
	}

	params := c.Request().URL.Query()
	if params.Encode() == "" && c.Request().Method == "POST" {
		err = c.Request().ParseForm()
		if err != nil {
			return goth.User{}, err
		}

		params = c.Request().Form
	}

	_, err = sess.Authorize(provider, params)
	if err != nil {
		return goth.User{}, err
	}

	gothUser, err := provider.FetchUser(sess)
	return gothUser, err
}

func (a *Auth) validateState(c echo.Context, s goth.Session) error {
	rawAuthURL, err := s.GetAuthURL()
	if err != nil {
		return err
	}

	authURL, err := url.Parse(rawAuthURL)
	if err != nil {
		return err
	}

	reqState := a.GetState(c)

	originalState := authURL.Query().Get("state")
	if originalState != "" && originalState != reqState {
		return errors.New("state token mismatch")
	}

	return nil
}

func (a *Auth) GetState(c echo.Context) string {
	return gothic.GetState(c.Request())
}

func (a *Auth) SetState(c echo.Context) string {
	return gothic.SetState(c.Request())
}

func (a *Auth) Logout(c echo.Context) error {
	return gothic.Logout(c.Response(), c.Request())
}

func (a *Auth) GetProviderName(c echo.Context) (string, error) {

	if p := c.Param("provider"); p != "" {
		return p, nil
	}

	return gothic.GetProviderName(c.Request())
}

func (a *Auth) StoreInSession(key string, value string, c echo.Context) error {
	return gothic.StoreInSession(key, value, c.Request(), c.Response())
}

func (a *Auth) GetFromSession(key string, c echo.Context) (string, error) {
	return gothic.GetFromSession(key, c.Request())
}
