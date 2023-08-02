package restapi

import (
	"encoding/json"
	"github.com/TJxiaobao/go-ddd-template/pkg/encode"
	"github.com/TJxiaobao/go-ddd-template/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Response struct {
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data,omitempty"`
	DataVersion string      `json:"data_version,omitempty"`
	RequestId   string      `json:"request_id,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	sendResponse(c, http.StatusOK, data, nil)
}

func Failed(c *gin.Context, err error) {
	sendResponse(c, http.StatusOK, nil, err)
}

func FailedWithStatus(c *gin.Context, err error, httpStatus int) {
	sendResponse(c, httpStatus, nil, err)
}

func sendResponse(c *gin.Context, httpStatus int, data interface{}, err error) {
	bizErr := errno.AssertBizError(err)
	// 传给其他middleware处理
	c.Set("x-bizError", bizErr)
	c.Set("x-httpStatus", httpStatus)
	// 将业务错误码放入response header中
	c.Writer.Header().Add("x-biz-code", strconv.Itoa(bizErr.Code()))
	// 返回json格式数据
	c.JSON(httpStatus, generateResponseDataWithVersion(c, bizErr, data))
}

func generateResponseDataWithVersion(c *gin.Context, err errno.BizError, data interface{}) *Response {
	resp := &Response{
		RequestId: GetRequestId(c),
		Code:      err.Code(),
		Message:   err.Message(),
		Data:      data,
	}

	if data == nil {
		return resp
	}

	// 计算data的hash并进行比较
	if c.Request.Header.Get("x-accept-version") == "crc32" {
		b, _ := json.Marshal(data)
		if len(b) > 0 {
			resp.DataVersion = encode.Crc32HashCode(b)
			c.Set("x-data-version", resp.DataVersion)
		}
		lastHash := c.Query("data_version")
		if lastHash != "" && lastHash == resp.DataVersion {
			resp.Data = nil
		}
	}

	return resp
}
