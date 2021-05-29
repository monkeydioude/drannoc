package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/pkg/entity"
	"github.com/monkeydioude/drannoc/pkg/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
	"github.com/monkeydioude/drannoc/pkg/service"
)

// AddNewTrade add a new trade in the history of a user
// POST /trade
func AddNewTrade(c *gin.Context) {
	userID := c.GetString(c.GetString("ConsumerLabel"))

	if userID == "" {
		res.Write(c, res.ServiceUnavailable("could not find userID", "no consumer in header"))
		return
	}

	trade := entity.NewTrade()
	// parse body and unmarshal a trade
	err := service.EntityFromRequestBody(c.Request.Body, trade)
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not add new trade", err.Error()))
		return
	}
	tradeRepo := repository.NewTrade()

	if ok, response := tradeRepo.ParentExists(trade); !ok && response != nil {
		res.Write(c, response)
		return
	}

	// remove ID if one was passed
	trade.ID = ""

	// force created_at and modified_at
	trade.Created_at = time.Now().UnixNano()
	trade.Modified_at = trade.Created_at

	trade.User_id = userID

	if !trade.IsStorable() {
		res.Write(c, res.BadRequest("incomplete payload"))
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

// EditTrade edits a trade using a trade_id
// PUT /trade/:trade_id
func EditTrade(c *gin.Context) {
	userID := c.GetString(c.GetString("ConsumerLabel"))

	if userID == "" {
		res.Write(c, res.ServiceUnavailable("could not find userID", "no consumer in header"))
		return
	}
	tradeRepo := repository.NewTrade()

	trade := entity.NewTrade()
	exist, err := tradeRepo.FindFirstByID(trade, c.Param("trade_id"))
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not edit trade", err.Error()))
		return
	}

	if exist == nil {
		res.Write(c, res.BadRequest("could not find trade_id"))
		return
	}

	if trade.User_id != userID {
		res.Write(c, res.BadRequest("user_id does not match or trade does not have any user_id"))
		return
	}

	tradeBody := entity.NewTrade()
	// parse body and unmarshal a trade
	err = service.EntityFromRequestBody(c.Request.Body, tradeBody)
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not edit trade", err.Error()))
		return
	}

	trade.UpdateWith(tradeBody)
	err = tradeRepo.Save(trade)
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not edit trade", err.Error()))
		return
	}

	res.Ok(c, gin.H{
		"data": trade.ID,
	})
	// fields, values := make(map[string]interface{})
}
