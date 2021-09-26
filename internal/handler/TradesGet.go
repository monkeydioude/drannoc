package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/repository"
	res "github.com/monkeydioude/drannoc/internal/response"
	"github.com/monkeydioude/drannoc/pkg/db"
)

// TradesGet retrieves every trade for a user
// Filters:
//	- coin_id string
func TradesGet(c *gin.Context) {
	userID := c.GetString(c.GetString("ConsumerLabel"))

	if userID == "" {
		res.Write(c, res.ServiceUnavailable("could not find userID", "no consumer in header"))
		return
	}

	query := db.NewQuery().WithFilters(db.KV("user_id", userID))
	coinID, ok := c.GetQuery("coin_id")
	if !ok {
		res.Write(c, res.BadRequest("coin_id must be provided in query params"))
		return
	}

	query = query.WithFilters(db.KV("coin_id", coinID))
	tradeRepo := repository.NewTrade()
	// retrieving cursor containing matching records
	cursor, errRes := tradeRepo.QueryAll(query)
	if errRes != nil {
		res.Write(c, errRes)
		return
	}

	// creating response arrays from cursor size
	results := make([]entity.Trade, cursor.RemainingBatchLength())
	// unmarshall results from mongo cursor
	cursor.All(tradeRepo.GetContext(), &results)

	for i := range results {
		results[i].User_id = ""
	}

	res.Ok(c, gin.H{
		"data": results,
	})
}
