package response

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// BadRequest = 400 response code
func BadRequest(c *gin.Context, msg string) {
	log.Info(msg)
	c.JSON(400, gin.H{
		"message": msg,
	})
}

// ServiceUnavailable = 503 response code
func ServiceUnavailable(c *gin.Context, msg string) {
	log.Error(msg)
	c.JSON(503, gin.H{
		"message": msg,
	})
}

// Ok = 200 response code
func Ok(c *gin.Context, res map[string]interface{}) {
	log.Info(res)
	c.JSON(200, res)
}
