package server

import (
	egoth "github.com/nabowler/echo-gothic"
)

func (s *Server) ProviderAuthCallbackHandler(c echo.Context) error {

	user, err := echogothic.CompleteUserAuth(c)
}
