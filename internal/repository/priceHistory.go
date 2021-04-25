package repository

import (
	"context"
	"time"

	"github.com/monkeydioude/drannoc/internal/db"
)

type PriceHistory struct {
	BaseRepo
}

func NewPriceHistory() *PriceHistory {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	return &PriceHistory{
		BaseRepo: BaseRepo{
			context:    ctx,
			cancelFunc: cancel,
			collection: db.Database(db.CoinsDbName).Collection("price_history"),
		},
	}
}
