package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	C_UNDERLINE byte   = '_'
	EMPTY       string = ""
	COLON       string = ":"
	COMMA       string = ","
	UNDERLINE   string = string(C_UNDERLINE)
	DASHED      string = "-"
	SLASH       string = "/"
	DOT         string = "."
	SPACE       string = " "
	QUESTION    string = "?"
	PIPE        string = "|"
)

// 单词全部转化为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 单词全部转化为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线单词转为大写驼峰单词
func ToCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	cases := cases.Title(language.English)
	s = cases.String(s)
	return strings.Replace(s, " ", "", -1)
}

// 驼峰单词转下划线单词
func ToUnderlineCase(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, rune(C_UNDERLINE))
			}
			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}
