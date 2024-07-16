package validatorx

import (
	"regexp"
	"strconv"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)

// CheckMaxStringLenWithoutNumberSuffix 检查去掉'-数字'尾字符串后的字符串长度是否小于等于指定长度 并且一个中文算1个字符
func CheckMaxStringLenWithoutNumberSuffix(checkString string, maxLen int) bool {
	suffixRegexStr := `-[\d]+$`
	suffixRegex, _ := regexp.Compile(suffixRegexStr)

	matchLoc := suffixRegex.FindStringIndex(checkString)
	if 0 == len(matchLoc) {
		return utf8.RuneCountInString(checkString) <= maxLen
	}

	return utf8.RuneCountInString(checkString[:matchLoc[0]]) <= maxLen
}

func checkMaxStringLenWithoutNumberSuffix(fl validator.FieldLevel) bool {
	maxLen, _ := strconv.Atoi(fl.Param())

	return CheckMaxStringLenWithoutNumberSuffix(fl.Field().String(), maxLen)
}
