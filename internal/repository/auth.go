package repository

import (
	"context"
	"time"

	"github.com/monkeydioude/drannoc/internal/db"
	entityInt "github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/pkg/entity"
)

// AuthRepository would be the implementation of the Repository interface
// for the User Collection
type AuthRepository struct {
	BaseRepo
}

// NewAuthRepository returns a pointer to a UserRepository instance
func NewAuthRepository() *AuthRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	return &AuthRepository{
		BaseRepo: BaseRepo{
			context:    ctx,
			cancelFunc: cancel,
			collection: db.Database(db.AuthDbName).Collection("authentication"),
		},
		// @TODO handle database name better
	}
}

// FindFirst calls the BaseRepo's FindFirst, adding the kind of entity
// it awaits.
func (repo *AuthRepository) FindFirst(filter Filter) (entity.Entity, error) {
	return repo.BaseRepo.FindFirst(&entityInt.Auth{}, filter)
}
