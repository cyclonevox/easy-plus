package validatorx

import "github.com/go-playground/validator/v10"

// CheckPassword 检查密码
// 密码规则
// 只支持数字、字母、特殊字符
// 每种类型至少一个
// 密码长度至少8位 在外边验证
func CheckPassword(password string /*, minLength, maxLength int*/) bool {
	var digit, letter, special bool

	for _, r := range []rune(password) {
		if r < 32 || (r > 126 && r < 256) {
			return false
		}

		if '0' <= r && '9' >= r {
			digit = true
		} else if 'a' <= r && 'z' >= r || 'A' <= r && 'Z' >= r {
			letter = true
		} else {
			special = true
		}

		if digit && letter && special {
			return true
		}
	}

	return false
}

func checkPassword(fl validator.FieldLevel) bool {
	return CheckPassword(fl.Field().String())
}
