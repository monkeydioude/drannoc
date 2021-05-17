package repository

import (
	"context"
	"time"

	"github.com/monkeydioude/drannoc/pkg/db"
	"github.com/monkeydioude/drannoc/pkg/entity"
)

// User would be the implementation of the Repository interface
// for the User Collection
type UserPreferences struct {
	BaseRepo
}

// NewUser returns a pointer to a User instance
func NewUserPreferences() *UserPreferences {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	return &UserPreferences{
		BaseRepo: BaseRepo{
			context:    ctx,
			cancelFunc: cancel,
			collection: db.Database(db.AuthDbName).Collection("user_preferences"),
		},
		// @TODO handle database name better
	}
}

// Load the document from the DB
func (repo *UserPreferences) Load(prefs *entity.UserPreferences, userID string) (entity.Entity, error) {
	return repo.FindFirst(prefs, db.Filter{"id": userID}, nil)
}
