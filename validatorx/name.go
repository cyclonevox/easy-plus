package validatorx

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// CheckName 只允许阳间名字 不允许含有^%&',;=?$
func CheckName(name string) bool {
	regular := `^[^\^%&',;=?$\x22]+$`
	reg := regexp.MustCompile(regular)

	return reg.MatchString(name)
}

func checkName(fl validator.FieldLevel) bool {
	return CheckName(fl.Field().String())
}

func ChineseOrEnglishOrNumbersOrSpace(name string) bool {
	regular := `^[A-Za-z0-9\p{Han}\s]*$`
	reg := regexp.MustCompile(regular)

	return reg.MatchString(name)
}

func chineseOrEnglishOrNumbersOrSpace(fl validator.FieldLevel) bool {
	return ChineseOrEnglishOrNumbersOrSpace(fl.Field().String())
}
