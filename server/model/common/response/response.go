package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    ResponseCode    `json:"code"`
	Data    interface{}     `json:"data"`
	Message ResponseMessage `json:"message"`
}

type ResponseCode int

const (
	Success ResponseCode = 0
	Error   ResponseCode = 1
)

type ResponseMessage string

const (
	SuccessMessage ResponseMessage = "success"
	ErrorMessage   ResponseMessage = "error"
)

func Result(code ResponseCode, data interface{}, msg ResponseMessage, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: msg,
	})
}
func Ok(c *gin.Context) {
	Result(Success, map[string]interface{}{}, SuccessMessage, c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, ResponseMessage(message), c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(Success, data, SuccessMessage, c)
}

func OkWithMessageAndData(message string, data interface{}, c *gin.Context) {
	Result(Success, data, ResponseMessage(message), c)
}

func Fail(c *gin.Context) {
	Result(Error, map[string]interface{}{}, ErrorMessage, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(Error, map[string]interface{}{}, ResponseMessage(message), c)
}

func FailWithMessageAndData(message string, data interface{}, c *gin.Context) {
	Result(Error, data, ResponseMessage(message), c)
}

func FailWithCodeAndMessage(code ResponseCode, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, ResponseMessage(message), c)
}

func FailWithCodeAndMessageAndData(code ResponseCode, message string, data interface{}, c *gin.Context) {
	Result(code, data, ResponseMessage(message), c)
}
