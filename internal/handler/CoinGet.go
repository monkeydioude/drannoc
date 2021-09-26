package handler

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/repository"
	res "github.com/monkeydioude/drannoc/internal/response"
	"github.com/monkeydioude/drannoc/pkg/db"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

func getOrder(c *gin.Context) int {
	order := -1
	if ord, ok := c.GetQuery("order"); ok {
		order, _ = strconv.Atoi(ord)
	}

	return order
}

func getDuration(c *gin.Context) int64 {
	var duration int64 = 3600000

	if dur, ok := c.GetQuery("duration"); ok {
		duration, _ = strconv.ParseInt(dur, 10, 64)
	}

	return duration
}

func filtersAddCreatedAt(c *gin.Context, filters db.Filter, duration int64) db.Filter {
	// milliseconds in db
	created_at := time.Now().UnixNano() / 1000000

	if ca, ok := c.GetQueryMap("created_at"); ok {
		for k, v := range ca {
			value, _ := strconv.ParseInt(v, 10, 64)
			filters.AddFilter("created_at", k, value)
		}
	} else {
		filters.Add("created_at", db.Filter{
			"$gte": created_at - duration,
		})
	}

	return filters
}

// CoinGet
// GET /coin/:coin_id
func CoinGet(c *gin.Context) {
	// filters used by mongodb.s Find
	filters := db.Filter{}
	// options used by mongodb.s Find
	options := db.NewOptions().Proj("created_at", 1)
	order := getOrder(c)
	duration := getDuration((c))

	filters, options.Projection = coinsFiltersProjects(
		c.Param("coin_id"),
		filters,
		options.Projection.(db.Filter),
	)
	filters = filtersAddCreatedAt(c, filters, duration)
	options.Sort = db.Filter{"created_at": order}

	repo := repository.NewPriceHistory()

	// triggers query to mongodb and retrieves a Cursor
	cursor, err := repo.GetCollection().Find(
		repo.GetContext(),
		filters,
		(*mongoOptions.FindOptions)(options),
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
