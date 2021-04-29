package repository

import (
	"context"
	"time"

	"github.com/monkeydioude/drannoc/internal/pkg/db"
	"github.com/monkeydioude/drannoc/internal/pkg/entity"
	iEntity "github.com/monkeydioude/drannoc/pkg/entity"
)

// UserID type
type UserID string

// Load retrieve a User document using only its ID
func (ID UserID) Load() (*entity.User, error) {
	repo := NewUser()
	user := &entity.User{
		ID: string(ID),
	}
	_, err := repo.Load(user)

	return user, err
}

// User would be the implementation of the Repository interface
// for the User Collection
type User struct {
	BaseRepo
}

// NewUser returns a pointer to a User instance
func NewUser() *User {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	return &User{
		BaseRepo: BaseRepo{
			context:    ctx,
			cancelFunc: cancel,
			collection: db.Database(db.AuthDbName).Collection("user"),
		},
		// @TODO handle database name better
	}
}

// Load the document from the DB
func (repo *User) Load(user *entity.User) (iEntity.Entity, error) {
	return repo.FindFirst(user, Filter{"id": user.ID})
}

// Load the document from the DB
func (repo *User) LoadFromCredentials(user *entity.User) (iEntity.Entity, error) {
	return repo.FindFirst(user, Filter{"login": user.Login, "password": user.Password})
}

// Create is like store but do some setup before
func (repo *User) Create(user *entity.User) (string, error) {
	user.Created = time.Now().Unix()
	return repo.Store(user)
}
