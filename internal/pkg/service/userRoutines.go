package service

import (
	"errors"
	"io"

	repo "github.com/monkeydioude/drannoc/internal/pkg/repository"

	"github.com/monkeydioude/drannoc/internal/pkg/entity"
)

// UserCreate is the service handling user creation.
// Creates a new user with encrypted password, stores the user
// and the auth in different buckets.
func UserCreate(
	body io.ReadCloser,
	userRepo *repo.User,
	tokenRepo *repo.AuthToken,
) (*entity.AuthToken, error) {
	user := &entity.User{}
	err := EntityFromRequestBody(body, user)

	if err != nil {
		return nil, err
	}
	user.PasswordEncrypt()

	// attempting to create a token
	token, err := CreateAuthTokenNow(tokenRepo, user.ID)
	if err != nil {
		return nil, err
	}

	// loading user to check if existing
	res, err := userRepo.LoadFromCredentials(user)

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
