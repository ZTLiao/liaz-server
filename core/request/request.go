package request

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHeaders(c *gin.Context) map[string]any {
	header := c.Request.Header
	var headerMap = make(map[string]any, len(header))
	for k, v := range header {
		headerMap[k] = v
	}
	return headerMap
}

func GetQueryParams(c *gin.Context) map[string]any {
	query := c.Request.URL.Query()
	var queryMap = make(map[string]any, len(query))
	for k := range query {
		queryMap[k] = c.Query(k)
	}
	return queryMap
}

func GetPostFormParams(c *gin.Context) (map[string]any, error) {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		if !errors.Is(err, http.ErrNotMultipart) {
			return nil, err
		}
	}
	var postMap = make(map[string]any, len(c.Request.PostForm))
	for k, v := range c.Request.PostForm {
		if len(v) > 1 {
			postMap[k] = v
		} else if len(v) == 1 {
			postMap[k] = v[0]
		}
	}
	return postMap, nil
}

func GetBodyParams(c *gin.Context) (string, error) {
	reqBytes, err := c.GetRawData()
	if err != nil {
		return "", err
	}
	// 请求包体写回
	if len(reqBytes) > 0 {
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBytes))
	}
	return string(reqBytes), nil
}
