package service

import (
	"github.com/monkeydioude/drannoc/internal/body"
	// "github.com/monkeydioude/drannoc/internal/bucket"
	"github.com/monkeydioude/drannoc/internal/entity"
	// "github.com/monkeydioude/drannoc/internal/misc"
)

// Authenticate service handles returning an authentification token
// from /auth endpoint
func Authenticate(loginData body.LoginData) (*entity.AuthToken, error) {
	// authBucket := bucket.Auth(nil)
	// authTokenBucket := bucket.AuthToken(nil)

	// u, err := authBucket.Get(loginData.Login)
	// if err != nil {
	// 	return nil, err
	// }

	// if u == nil {
	// 	return nil, errors.New("Could not authenticate using login data")
	// }

	// encPasswd := entity.NewAuth(loginData.Login, loginData.Password).GetPassword()
	// if string(u) != encPasswd {
	// 	return nil, err
	// }

	// token := entity.NewAuthToken(encPasswd, time.Now(), misc.TokenDuration)
	// _, err = authTokenBucket.Store(token)
	// if err != nil {
	// 	return nil, err
	// }

	// return token, nil
	return nil, nil
}
