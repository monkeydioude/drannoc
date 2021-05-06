package repository

import (
	"context"
	"time"

	"github.com/monkeydioude/drannoc/pkg/db"
	"github.com/monkeydioude/drannoc/pkg/entity"
)

// CoinInfo is a repository for all coin data related
type CoinInfo struct {
	BaseRepo
}

// NewAuthRepository returns a pointer to a User instance
func NewCoinInfo() *CoinInfo {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	return &CoinInfo{
		BaseRepo: BaseRepo{
			context:    ctx,
			cancelFunc: cancel,
			collection: db.Database(db.CoinsDbName).Collection("coin_info"),
		},
		// @TODO handle database name better
	}
}

func (repo *CoinInfo) LoadAll() ([]*entity.CoinInfo, error) {
	arr := []*entity.CoinInfo{}
	cursor, err := repo.GetCollection().Find(
		repo.GetContext(),
		Filter{},
		nil,
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
