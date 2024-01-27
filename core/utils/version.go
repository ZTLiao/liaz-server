package utils

import (
	"strconv"
	"strings"
)

func ParseAppVersion(appVersion string) (int64, error) {
	return strconv.ParseInt(strings.ReplaceAll(appVersion, DOT, "0"), 10, 64)
}

func CompareAppVersion(appVersion1 string, appVersion2 string) (int, error) {
	version1, err := ParseAppVersion(appVersion1)
	if err != nil {
		return 0, err
	}
	version2, err := ParseAppVersion(appVersion2)
	if err != nil {
		return 0, err
	}
	if version1 > version2 {
		return 1, nil
	} else if version1 == version2 {
		return 0, nil
	} else if version1 < version2 {
		return -1, nil
	}
	return 0, nil
}
