package handler

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/pkg/db"
	"github.com/monkeydioude/drannoc/pkg/entity"
	"github.com/monkeydioude/drannoc/pkg/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

func getCoins(c *gin.Context, proj interface{}) (db.Filter, interface{}) {
	filters := db.Filter{}
	if coins, ok := c.GetQuery("coins"); ok && coins != "" {
		filters, proj = coinsFiltersProjects(
			coins,
			filters,
			proj.(db.Filter),
		)
	}

	return filters, proj
}

// coinsFiltersProjects generates Find Filters and Find Options
// using a string list of coins
func coinsFiltersProjects(
	coins string,
	filters,
	projections db.Filter,
) (db.Filter, db.Filter) {
	or := []db.Filter{}

	for _, v := range strings.Split(coins, ",") {
		key := "coins." + v
		f := db.Filter{}
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
	// milliseconds in db
	created_at := time.Now().UnixNano() / 1000000

	// options used by mongodb.s Find
	options := db.NewOptions().Proj("created_at", 1)
	order := getOrder(c)
	duration := getDuration((c))

	// filters used by mongodb.s Find
	filters, proj := getCoins(c, options.Projection)
	options.Projection = proj

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

	repo := repository.NewPriceHistory()

	options.Sort = db.Filter{"created_at": order}

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
