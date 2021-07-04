package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/pkg/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

// TradeDelete delete a trade using a trade_id
// DELETE /trade/:trade_id
func TradeDelete(c *gin.Context) {
	userID := c.GetString(c.GetString("ConsumerLabel"))

	if userID == "" {
		res.Write(c, res.BadRequest("could not find userID"))
		return
	}

	tradeID := c.Param("trade_id")
	err := repository.NewTrade().DeleteWithRelatives(tradeID)

	if err != nil {
		res.Write(c, res.ServiceUnavailable(fmt.Sprintf("could not delete trade with id %s", tradeID), "could not delete trade"))
		return
	}

	res.Ok(c, gin.H{
		"data": "ok",
	})
}
