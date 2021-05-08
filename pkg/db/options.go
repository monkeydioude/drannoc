package db

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOptions options.FindOptions

func NewOptions() *FindOptions {
	return &FindOptions{
		Projection: Filter{},
	}
}

// Proj add key value to the Projection filters
func (fo *FindOptions) Proj(name string, value interface{}) *FindOptions {
	proj := fo.Projection.(Filter)
	proj[name] = value
	fo.Projection = proj
	return fo
}
