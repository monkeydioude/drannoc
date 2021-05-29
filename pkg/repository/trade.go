package repository

import (
	"context"
	"time"

	"github.com/monkeydioude/drannoc/pkg/db"
	"github.com/monkeydioude/drannoc/pkg/entity"
	"github.com/monkeydioude/drannoc/pkg/response"
)

// User would be the implementation of the Repository interface
// for the User Collection
type Trade struct {
	BaseRepo
}

// NewUser returns a pointer to a User instance
func NewTrade() *Trade {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	return &Trade{
		BaseRepo: BaseRepo{
			context:    ctx,
			cancelFunc: cancel,
			collection: db.Database(db.AuthDbName).Collection("trade"),
		},
		// @TODO handle database name better
	}
}

// ParentExists tries to verifies that transaction has a parent
func (t *Trade) ParentExists(e *entity.Trade) (bool, *response.Response) {
	// no parent_id, nothing to do
	if e.Parent_id == "" {
		return false, nil
	}

	parent := entity.NewTrade()
	result, err := t.FindFirstByID(parent, e.Parent_id)
	// Error while fetching parent
	if err != nil {
		return false, response.ServiceUnavailable("error while fetching parent_id", err.Error())
	}

	// Could not find parent_id
	if result == nil || result.GetID() == "" {
		return false, response.BadRequest("parent_id not found")
	}

	return true, nil
}
