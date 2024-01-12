package utils

import "strings"

func In(column string, objs ...interface{}) (string, []interface{}) {
	var builder strings.Builder
	var params = make([]interface{}, 0)
	builder.WriteString(column + " in (")
	for i, length := 0, len(objs); i < length; i++ {
		builder.WriteString(QUESTION)
		params = append(params, objs[i])
		if i != length-1 {
			builder.WriteString(COMMA)
		}
	}
	builder.WriteString(")")
	return builder.String(), params
}
