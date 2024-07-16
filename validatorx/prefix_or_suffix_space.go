package validatorx

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// 检查前后缀空白字符
func CheckPrefixOrSuffixSpace(str string) bool {
	trimStr := strings.TrimSpace(str)

	return len(str) == len(trimStr)
}

// 检查前后缀空白字符
func checkPrefixOrSuffixSpace(fl validator.FieldLevel) bool {
	return CheckPrefixOrSuffixSpace(fl.Field().String())
}
