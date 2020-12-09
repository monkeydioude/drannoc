package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	log "github.com/sirupsen/logrus"
)

// Redirect = 301
func Redirect(c *gin.Context, url string) {
	c.JSON(301, gin.H{
		"url": url,
	})
}

// BadRequest = 400 response code
func BadRequest(c *gin.Context, msg string) {
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
	data, _ := json.Marshal(res)
	c.Render(http.StatusOK, render.Data{
		ContentType: "application/json",
		Data:        data,
	})
}
