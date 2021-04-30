package repository

import (
	"context"
	"time"

	"github.com/monkeydioude/drannoc/pkg/db"
	"github.com/monkeydioude/drannoc/pkg/entity"
	iEntity "github.com/monkeydioude/drannoc/pkg/entity"
)

// AuthRepository would be the implementation of the Repository interface
// for the User Collection
type AuthToken struct {
	BaseRepo
}

// NewAuthRepository returns a pointer to a User instance
func NewAuthToken() *AuthToken {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	return &AuthToken{
		BaseRepo: BaseRepo{
			context:    ctx,
			cancelFunc: cancel,
			collection: db.Database(db.AuthDbName).Collection("authToken"),
		},
		// @TODO handle database name better
	}
}

func (repo *AuthToken) Load(token *entity.AuthToken) (iEntity.Entity, error) {
	return repo.FindFirst(token, Filter{"token": token.Token})
}
