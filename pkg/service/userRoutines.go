package service

import (
	"errors"
	"io"

	repo "github.com/monkeydioude/drannoc/pkg/repository"

	"github.com/monkeydioude/drannoc/pkg/entity"
)

// UserCreate is the service handling user creation.
// Creates a new user with encrypted password, stores the user
// and the auth in different buckets.
func UserCreate(
	body io.ReadCloser,
	userRepo *repo.User,
	tokenRepo *repo.AuthToken,
) error {
	user := &entity.User{}
	err := EntityFromRequestBody(body, user)

	if err != nil {
		return err
	}
	user.PasswordEncrypt()

	// loading user to check if existing
	res, err := userRepo.LoadFromCredentials(user)
	// sum' happened
	if err != nil {
		return err
	}
	// user exists
	if res != nil {
		return errors.New("login already exists")
	}

	_, err = userRepo.Create(user)
	// storing the user in collection
	if err != nil {
		return err
	}
	return nil
}
