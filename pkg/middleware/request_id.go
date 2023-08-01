package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

const (
	requestIdKey = "request_id"
)

func RequestId(c *gin.Context) {
	// Check for incoming header, use it if exists
	requestId := c.Request.Header.Get(requestIdKey)

	// Create request id with UUID4
	if requestId == "" {
		u4, _ := uuid.NewV4()
		requestId = u4.String()
	}

	// Expose it for use in the application
	c.Set(requestIdKey, requestId)

	// Set X-Request-Id header
	c.Writer.Header().Set(requestIdKey, requestId)
	c.Next()
}
