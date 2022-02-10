package gin

import "github.com/gin-gonic/gin"

func GetUserIDFromContext(c *gin.Context) string {
	return c.GetString(c.GetString("ConsumerLabel"))
}
