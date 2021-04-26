package handler

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// coinsFiltersProjects generates Find Filters and Find Options
// using a string list of coins
func coinsFiltersProjects(
	coins string, filters,
	projections repository.Filter,
) (repository.Filter, repository.Filter) {
	or := []repository.Filter{}

	for _, v := range strings.Split(coins, ",") {
		key := "coins." + v
		f := repository.Filter{}
		f.AddFilter(key, "$exists", true)
		or = append(or, f)
		projections.Add(key, 1)
	}

	filters["$or"] = or

	return filters, projections
}

// CoinsGet retrieves coins rate history using filters.
// Filters:
//	- order int
// 	- duration int64
//  - coins string (comma "," separated list of coins)
// 	- created_at map[string]int64 (using mongodb operators as keys)
//
// GET /coins
func CoinsGet(c *gin.Context) {
	// order defined in with
	order := -1
	// milliseconds in db
	created_at := time.Now().UnixNano() / 1000000
	// filters used by mongodb.s Find
	filters := repository.Filter{}
	// options used by mongodb.s Find
	options := &options.FindOptions{
		Projection: repository.Filter{
			"created_at": 1,
		},
	}

	var duration int64 = 3600000

	if ord, ok := c.GetQuery("order"); ok {
		order, _ = strconv.Atoi(ord)
	}

	if dur, ok := c.GetQuery("duration"); ok {
		duration, _ = strconv.ParseInt(dur, 10, 64)
	}

	if coins, ok := c.GetQuery("coins"); ok && coins != "" {
		filters, options.Projection = coinsFiltersProjects(
			coins,
			filters,
			options.Projection.(repository.Filter),
		)
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

	options.Sort = repository.Filter{"created_at": order}

	// triggers query to mongodb and retrieves a Cursor
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
