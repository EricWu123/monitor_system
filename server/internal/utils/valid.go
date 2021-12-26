package utils

import (
	"regexp"
)

// 白名单校验字符串
func CheckStrWhite(str string, pattern string, maxLen int) (bool, error) {
	if maxLen > 0 {
		if len(str) > maxLen {
			return false, nil
		}
	}
	matched, err := regexp.MatchString(pattern, str)
	return matched, err
}

// 黑名单校验字符串
func CheckStrBlack(str string, pattern string, maxLen int) (bool, error) {
	if maxLen > 0 {
		if len(str) > maxLen {
			return false, nil
		}
	}
	matched, err := regexp.MatchString(pattern, str)
	return !matched, err
}
