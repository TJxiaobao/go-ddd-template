package middleware

import (
	"github.com/TJxiaobao/go-ddd-template/pkg/errno"
	"github.com/TJxiaobao/go-ddd-template/pkg/restapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFound(c *gin.Context) {
	restapi.FailedWithStatus(c, errno.NewSimpleError(errno.ErrRouteNotFound, nil), http.StatusNotFound)
}
