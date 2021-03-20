package service

import (
	"errors"

	"github.com/monkeydioude/drannoc/internal/body"
	repo "github.com/monkeydioude/drannoc/internal/repository"

	"github.com/monkeydioude/drannoc/internal/entity"
)

// UserCreate is the service handling user creation.
// Creates a new user with encrypted password, stores the user
// and the auth in different buckets.
func UserCreate(
	loginData body.LoginData,
	userRepo *repo.User,
	tokenRepo *repo.AuthToken,
) (*entity.AuthToken, error) {

	if len(loginData.Login) == 0 || len(loginData.Password) == 0 {
		return nil, errors.New("a login and a password must be given")
	}

	// attempting to create a token
	token, err := CreateAuthTokenNow(tokenRepo)
	if err != nil {
		return nil, err
	}

	user, err := entity.NewUser(loginData.Login, loginData.Password)
	// could not verify user
	if err != nil {
		return nil, err
	}
	// loading user to check if existing
	res, err := userRepo.Load(user)

	// sum' happened
	if err != nil {
		tokenRepo.Delete(token)
		return nil, err
	}
	// user exists
	if res != nil {
		tokenRepo.Delete(token)
		return nil, errors.New("login already exists")
	}

	_, err = userRepo.Create(user)
	// storing the user in collection
	if err != nil {
		tokenRepo.Delete(token)
		return nil, err
	}
	return token, nil
}
