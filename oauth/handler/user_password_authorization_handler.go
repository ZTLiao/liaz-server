package handler

import (
	"basic/storage"
	"context"
	"strconv"

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
	account, err := e.AccountDb.SignIn(username, password)
	if err != nil {
		return "", err
	}
	if account == nil {
		err = errors.ErrAccessDenied
		return
	}
	userID = strconv.FormatInt(account.UserId, 10)
	return
}
