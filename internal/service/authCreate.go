package service

import (
	"errors"

	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/repository"
)

// AuthCreate is the auth creation routine.
func AuthCreate(userID, password string) (*entity.Auth, error) {
	auth := entity.NewAuth(userID, password)
	authRepo := repository.NewAuthRepository()

	res, err := authRepo.FindFirst(repository.Filter{"userID": userID})
	if err != nil {
		return nil, err
	}
	if res != nil {
		return nil, errors.New("UserID already exists in auth")
	}

	_, err = authRepo.Store(auth)
	if err != nil {
		return nil, err
	}
	return auth, nil
}
