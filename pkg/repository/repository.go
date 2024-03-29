package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/monkeydioude/drannoc/pkg/db"
	"github.com/monkeydioude/drannoc/pkg/entity"
	res "github.com/monkeydioude/drannoc/pkg/response"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository is the base interface for DB requesting
type Repository interface {
	GetCollection() *mongo.Collection
	GetContext() context.Context
	CancelContext()
	Store(entity entity.Entity) error
	Load(entity.Entity) (entity.Entity, error)
	Delete(entity.Entity) error
}

// BaseRepo holds base generic repository functions.
// Like context handling methods
type BaseRepo struct {
	context    context.Context
	cancelFunc context.CancelFunc
	collection *mongo.Collection
}

// GetContext retrieves the context setup by the repository.
// Use when requesting. Implements Repository interface
func (repo BaseRepo) GetContext() context.Context {
	return repo.context
}

// CancelContext retrieves the cancelFunc generated along context.
// Used to cancel a context, thus a query. Implements Repository interface
func (repo BaseRepo) CancelContext() {
	repo.cancelFunc()
}

// GetCollection returns the collection object targeted,
// by a repository. Implements Repository interface.
func (repo BaseRepo) GetCollection() *mongo.Collection {
	return repo.collection
}

// QueryAll queries mongoDB using a db.Query struct
func (repo BaseRepo) QueryAll(q db.Query) (*mongo.Cursor, *res.Response) {
	cursor, err := repo.GetCollection().Find(
		repo.GetContext(),
		q.Filters,
		(*options.FindOptions)(&q.Options),
	)

	if err != nil {
		return nil, res.ServiceUnavailable("could not query db", err.Error())
	}

	return cursor, nil
}

// FindFirst finds the first element using a filter
func (repo BaseRepo) FindFirst(
	entity entity.Entity,
	filter db.Filter,
	projection db.Filter,
) (entity.Entity, error) {
	err := repo.GetCollection().FindOne(
		repo.GetContext(),
		filter,
		&options.FindOneOptions{
			Projection: projection,
		},
	).Decode(entity)

	if entity.GetID() == "" {
		return nil, nil
	}

	return entity, err
}

// FindFirstByID returns the first element using its ID
func (repo BaseRepo) FindFirstByID(entity entity.Entity, id string) (entity.Entity, error) {
	return repo.FindFirst(entity, db.Filter{"id": id}, nil)
}

// Store implements Repository interface. Store creates
// fill the paramatered passed entity with an ID
// then store it into DB
func (repo BaseRepo) Store(entity entity.Entity) (string, error) {
	entity.SetID(uuid.New().String())
	_, err := repo.GetCollection().InsertOne(repo.GetContext(), entity)

	if err != nil {
		return "", err
	}

	return entity.GetID(), nil
}

// Delete deletes multipe elements matching the document ID
func (repo BaseRepo) Delete(id string) error {
	_, err := repo.GetCollection().DeleteMany(repo.context, db.Filter{"id": id})
	return err
}

func (repo BaseRepo) Save(e entity.Entity) error {
	upsert := true
	options := &options.UpdateOptions{
		Upsert: &upsert,
	}
	update := map[string]entity.Entity{
		"$set": e,
	}
	_, err := repo.GetCollection().UpdateOne(repo.context, db.Filter{"id": e.GetID()}, update, options)
	return err
}

func (repo BaseRepo) Update(e entity.Entity, fields []string, values []interface{}) error {
	if len(fields) != len(values) {
		return errors.New("should have the same amount of fields and values")
	}
	upsert := true
	options := &options.UpdateOptions{
		Upsert: &upsert,
	}
	update := make(map[string]map[string]interface{})
	update["$set"] = make(map[string]interface{})
	for k, v := range fields {
		update["$set"][v] = values[k]
	}
	_, err := repo.GetCollection().UpdateOne(repo.context, db.Filter{"id": e.GetID()}, update, options)
	return err
}

func (repo BaseRepo) Find(
	ent entity.Entity,
	filters db.Filter,
	options *options.FindOptions,
	arr []entity.Entity,
) ([]entity.Entity, error) {
	cursor, err := repo.GetCollection().Find(
		repo.GetContext(),
		filters,
		options,
	)

	if err != nil {
		return nil, err
	}

	cursor.All(repo.GetContext(), &arr)

	if err != nil {
		return nil, err
	}

	return arr, nil
}
