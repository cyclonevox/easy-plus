package validatorx

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// CheckChinese 检查是否全中文
func CheckChinese(text string) bool {
	regular := `^[\p{Han}]*$`
	reg := regexp.MustCompile(regular)

	return reg.MatchString(text)
}

func checkChinese(fl validator.FieldLevel) bool {
	return CheckChinese(fl.Field().String())
}
