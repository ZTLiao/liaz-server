package handler

import (
	"basic/storage"
	"context"

	"github.com/go-oauth2/oauth2/v4/errors"
)

type UserPasswordAuthorizationHandler struct {
	ClientId  string
	AccountDb *storage.AccountDb
}

func (e *UserPasswordAuthorizationHandler) Authorize(ctx context.Context, clientID, username, password string) (userID string, err error) {
	if e.ClientId != clientID {
		err = errors.ErrUnauthorizedClient
		return
	}
	ok, err := e.AccountDb.IsExistForUsername(username, password)
	if err != nil {
		return "", err
	}
	if !ok {
		err = errors.ErrAccessDenied
		return
	}
	userID = username
	return
}
