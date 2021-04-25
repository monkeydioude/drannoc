package handler

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CoinsGet(c *gin.Context) {
	order := -1
	// milliseconds in db
	created_at := time.Now().UnixNano() / 1000000
	filters := repository.Filter{}
	var duration int64 = 3600000

	if ord, ok := c.GetQuery("order"); ok {
		order, _ = strconv.Atoi(ord)
	}

	if dur, ok := c.GetQuery("duration"); ok {
		duration, _ = strconv.ParseInt(dur, 10, 64)
	}

	if ca, ok := c.GetQueryMap("created_at"); ok {
		for k, v := range ca {
			value, _ := strconv.ParseInt(v, 10, 64)
			filters.AddFilter("created_at", k, value)
		}
	} else {
		filters.Add("created_at", repository.Filter{
			"$gte": created_at - duration,
		})
	}

	repo := repository.NewPriceHistory()

	options := &options.FindOptions{
		Sort: repository.Filter{"created_at": order},
	}

	cursor, err := repo.GetCollection().Find(
		repo.GetContext(),
		filters,
		options,
	)

	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not get coins data [1]", err.Error()))
		return
	}

	coinsArr := []entity.Stack{}
	cursor.All(repo.GetContext(), &coinsArr)

	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not get coins data [2]", err.Error()))
		return
	}
	res.Ok(c, gin.H{
		"data": coinsArr,
	})
}
