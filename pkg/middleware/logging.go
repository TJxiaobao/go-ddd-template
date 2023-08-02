package middleware

import (
	"bytes"
	"github.com/TJxiaobao/go-ddd-template/pkg/errno"
	util "github.com/TJxiaobao/go-ddd-template/pkg/iputil"
	"github.com/TJxiaobao/go-ddd-template/pkg/restapi"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		// Skip for the health check requests.
		if strings.HasPrefix(path, "/health") || strings.HasPrefix(path, "/swagger") || path == "/metrics" {
			return
		}

		// Read the Body content
		var (
			bodyBytes []byte
		)
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		apiLogger := log.WithFields(log.Fields{
			"log_type":     "HTTP_ACCESS",
			"request_time": time.Now().Format("2006-01-02 15:04:05"),
		})

		// Continue.
		c.Next()

		// Calculates the latency (µs).
		latency := time.Since(start) / 1000

		var (
			code    = errno.OK.Code
			message = errno.OK.Message
			errMsg  = ""
			header  = make(map[string]string, 0)
		)

		if err, ok := c.Get("x-bizError"); ok {
			bizErr := err.(errno.Error)
			code = bizErr.Code()
			message = bizErr.Message()
			if !bizErr.IsSuccess() {
				errMsg = bizErr.Error()
			}
		}

		if v, ok := c.Get("x-data-version"); ok {
			apiLogger.WithFields(log.Fields{"version": v})
		}

		apiLogger.WithFields(log.Fields{
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"header":      header,
			"query":       c.Request.URL.Query(),
			"success":     code == errno.OK.Code,
			"ret_code":    code,
			"ret_message": message,
			"latency":     latency,
			"biz_error":   errMsg,
			"remote":      util.GetRemoteRealIp(c),
			"request_id":  restapi.GetRequestId(c),
		}).Infof("")
	}
}
