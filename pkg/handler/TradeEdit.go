package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/pkg/entity"
	"github.com/monkeydioude/drannoc/pkg/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
	"github.com/monkeydioude/drannoc/pkg/service"
)

// TradeEdit edits a trade using a trade_id
// PUT /trade/:trade_id
func TradeEdit(c *gin.Context) {
	userID := c.GetString(c.GetString("ConsumerLabel"))

	if userID == "" {
		res.Write(c, res.BadRequest("could not find userID"))
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
}
