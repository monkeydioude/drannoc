package service

import (
	"errors"
	"time"

	"github.com/monkeydioude/drannoc/internal/body"
	"github.com/monkeydioude/drannoc/internal/bucket"
	"github.com/monkeydioude/drannoc/internal/misc"

	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/repository"
)

// UserCreate is the service handling user creation.
// Creates a new user with encrypted password, stores the user
// and the auth in different buckets.
func UserCreate(loginData body.LoginData) (*entity.AuthToken, error) {
	userEntity := entity.NewUser(loginData.Login)
	userRepo := repository.NewUserRepository()

	res, err := userRepo.FindFirst(repository.Filter{"login": userEntity.Login})
	if err != nil {
		return nil, err
	}
	if res != nil {
		return nil, errors.New("Login already exists in user")
	}

	auth, err := AuthCreate(userEntity.ID, loginData.Password)
	if err != nil {
		return nil, err
	}

	token, err := AuthTokenCreate(auth.GetPassword(), time.Now(), misc.TokenDuration)
	if err != nil {
		repository.NewAuthRepository().Delete(auth)
		return nil, err
	}

	_, err = userRepo.Store(userEntity)
	if err != nil {
		repository.NewAuthRepository().Delete(auth)
		bucket.AuthToken(nil).Delete(token.GetToken())
		return nil, err
	}
	return token, nil
}
