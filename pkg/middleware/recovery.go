package middleware

import (
	"fmt"
	"github.com/TJxiaobao/go-ddd-template/pkg/errno"
	"github.com/TJxiaobao/go-ddd-template/pkg/restapi"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		var err error
		if e, ok := recovered.(error); ok {
			err = e
		}
		if e, ok := recovered.(string); ok {
			err = errors.New(e)
		}
		stack := debug.Stack()
		log.Errorf("[PANIC][HTTP] request %v panic: %v, stack: %s", c.Request.URL, err, string(stack))
		err = fmt.Errorf("panic occurred, error: %v, stack: %s", err, stack)
		restapi.Failed(c, errno.NewError(errno.ErrInternalServerPanic, err))
	})
}
