package response

import (
	"net/http"
	"time"
)

type Response struct {
	Code      int         `json:"code"`           // 响应编码
	Message   string      `json:"message"`        // 响应信息
	Data      interface{} `json:"data,omitempty"` // 响应体
	Timestamp int64       `json:"timestamp"`      // 当前时间戳
}

// constructor
func New(code int, message string, data interface{}) *Response {
	response := new(Response)
	response.Code = code
	response.Message = message
	response.Data = data
	response.Timestamp = time.Now().UnixMilli()
	return response
}

// success
func Success() *Response {
	return ReturnOK(nil)
}

// fail
func Fail(message string) *Response {
	var code = http.StatusInternalServerError
	if len(message) == 0 {
		message = http.StatusText(code)
	}
	return New(code, message, nil)
}

// return ok
func ReturnOK(data interface{}) *Response {
	var code = http.StatusOK
	var message = http.StatusText(code)
	return New(code, message, data)
}

// return error
func ReturnError(code int, message string) *Response {
	return New(code, message, nil)
}
