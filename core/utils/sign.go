package utils

import (
	"core/logger"
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strconv"
	"strings"
)

func GenerateSign(params map[string][]string, timestamp int64, key string) string {
	if timestamp > 0 {
		var list = make([]string, 0)
		list = append(list, strconv.FormatInt(timestamp, 10))
		params["timestamp"] = list
	}
	var keys = make([]string, 0)
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var sb strings.Builder
	var index, length = 0, len(params)
	for _, key := range keys {
		value := params[key]
		sort.Strings(value)
		sb.Write([]byte(key))
		sb.Write([]byte("="))
		sb.Write([]byte(strings.Join(value, ",")))
		if index != (length - 1) {
			sb.Write([]byte("&"))
		}
		index++
	}
	sb.Write([]byte("&key="))
	sb.Write([]byte(key))
	logger.Info("sign : %s", sb.String())
	var md5New = md5.New()
	md5New.Write([]byte(sb.String()))
	return strings.ToUpper(hex.EncodeToString(md5New.Sum(nil)))
}
