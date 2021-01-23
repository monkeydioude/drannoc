package repository

import (
	"context"
	"time"

	"github.com/monkeydioude/drannoc/internal/db"
	entityInt "github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/pkg/entity"
)

// UserRepository would be the implementation of the Repository interface
// for the User Collection
type UserRepository struct {
	BaseRepo
}

// NewUserRepository returns a pointer to a UserRepository instance
func NewUserRepository() *UserRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	return &UserRepository{
		BaseRepo: BaseRepo{
			context:    ctx,
			cancelFunc: cancel,
			collection: db.Database(db.AuthDbName).Collection("user"),
		},
		// @TODO handle database name better
	}
}

// FindFirst calls the BaseRepo's FindFirst, adding the kind of entity
// it awaits.
func (repo *UserRepository) FindFirst(filter Filter) (entity.Entity, error) {
	return repo.BaseRepo.FindFirst(&entityInt.User{}, filter)
}
