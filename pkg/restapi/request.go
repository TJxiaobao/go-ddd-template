package restapi

import "github.com/gin-gonic/gin"

const (
	requestIdKey = "request_id"
)

func GetRequestId(c *gin.Context) string {
	v, ok := c.Get(requestIdKey)
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}
