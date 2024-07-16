package validatorx

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// CheckEndsWithIn 检查是否包含指定后缀列表中的一个
func CheckEndsWithIn(str string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}

	return false
}

func checkEndsWithIn(fl validator.FieldLevel) bool {
	params := strings.Split(fl.Param(), " ")
	for i := range params {
		params[i] = strings.TrimPrefix(params[i], "'")
		params[i] = strings.TrimSuffix(params[i], "'")
	}

	return CheckEndsWithIn(fl.Field().String(), params)
}
