package handler

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/repository"
	res "github.com/monkeydioude/drannoc/internal/response"
	"github.com/monkeydioude/drannoc/internal/service"
)

// TradeAdd add a new trade in the history of a user
// POST /trade
func TradeAdd(c *gin.Context) {
	userID := c.GetString(c.GetString("ConsumerLabel"))

	if userID == "" {
		res.Write(c, res.BadRequest("could not find userID"))
		return
	}

	trade := entity.NewTrade()
	// parse body and unmarshal a trade
	err := service.EntityFromRequestBody(c.Request.Body, trade)
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not add new trade", err.Error()))
		return
	}

	// case insensitive payload field
	trade.Direction = strings.ToUpper(trade.Direction)
	// erase ID, in case one was provided
	trade.ID = ""

	// force created_at and modified_at
	trade.Created_at = time.Now().UnixNano()
	trade.Modified_at = trade.Created_at

	trade.User_id = userID

	if !trade.IsStorable() {
		res.Write(c, res.BadRequest("incomplete payload"))
		return
	}

	tradeRepo := repository.NewTrade()

	if ok, response := tradeRepo.ParentExists(trade); !ok && response != nil {
		res.Write(c, response)
		return
	}

	id, err := tradeRepo.Store(trade)
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not add new trade", err.Error()))
		return
	}

	res.Ok(c, gin.H{
		"data": id,
	})
}
