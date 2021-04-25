package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/monkeydioude/drannoc/pkg/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Filter is as shortcut to map[string]interface{} (aka json)
type Filter map[string]interface{}

func (f Filter) Add(k string, v interface{}) {
	f[k] = v
}

func (f Filter) AddFilter(k, fk string, v interface{}) {
	filter := Filter{}

	if _, ok := f[k]; ok {
		switch f[k].(type) {
		case Filter:
			filter = f[k].(Filter)
		}
	}
	filter[fk] = v
	f.Add(k, filter)
}

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

// FindFirst finds the first element using a filter
func (repo BaseRepo) FindFirst(entity entity.Entity, filter Filter) (entity.Entity, error) {
	err := repo.GetCollection().FindOne(repo.GetContext(), filter).Decode(entity)

	if entity.GetID() == "" {
		return nil, nil
	}

	return entity, err
}

// FindFirstByID returns the first element using its ID
func (repo BaseRepo) FindFirstByID(entity entity.Entity, id string) (entity.Entity, error) {
	return repo.FindFirst(entity, Filter{"id": id})
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
func (repo BaseRepo) Delete(entity entity.Entity) error {
	_, err := repo.GetCollection().DeleteMany(repo.context, Filter{"id": entity.GetID()})
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
	_, err := repo.GetCollection().UpdateOne(repo.context, Filter{"id": e.GetID()}, update, options)
	return err
}

func (repo BaseRepo) Find(
	ent entity.Entity,
	filters Filter,
	options *options.FindOptions,
	arrBuilder func(len int) []entity.Entity,
) ([]entity.Entity, error) {
	cursor, err := repo.GetCollection().Find(
		repo.GetContext(),
		filters,
		options,
	)

	if err != nil {
		return nil, err
	}

	raw, err := cursor.Current.Elements()

	if err != nil {
		return nil, err
	}

	data := arrBuilder(len(raw))

	for cursor.Next(repo.GetContext()) {
		err = cursor.Decode(&ent)
		data = append(data, ent)
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}
