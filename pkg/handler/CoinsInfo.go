package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeydioude/drannoc/pkg/repository"
	res "github.com/monkeydioude/drannoc/pkg/response"
)

func CoinsInfo(c *gin.Context) {
	coinInfos, err := repository.NewCoinInfo().LoadAll()
	// error retrieving
	if err != nil {
		res.Write(c, res.ServiceUnavailable("could not retrieve coins info", err.Error()))
		return
	}

	res.Ok(c, gin.H{
		"data": coinInfos,
	})
}
