package validatorx

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func initTranslation(validate *validator.Validate, chinese ut.Translator) (err error) {
	if err = validate.RegisterTranslation(
		"max_len_without_number_suffix",
		chinese,
		func(ut ut.Translator) error {
			return ut.Add("max_len_without_number_suffix", "{0}去掉后缀后长度必须小于等于{1}", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("max_len_without_number_suffix", fe.Field(), fe.Param())
			return t
		},
	); nil != err {
		return
	}

	if err = validate.RegisterTranslation(
		"password",
		chinese,
		func(ut ut.Translator) error {
			return ut.Add("password", "密码只能由大小写字母、数字，特殊字符组成，且每种类型至少一位", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("password", fe.Field(), fe.Param())
			return t
		},
	); nil != err {
		return
	}

	if err = validate.RegisterTranslation(
		"chinese",
		chinese,
		func(ut ut.Translator) error {
			return ut.Add("chinese", "只能输入中文字符", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("chinese", fe.Field(), fe.Param())
			return t
		},
	); nil != err {
		return
	}

	if err = validate.RegisterTranslation(
		"sortby",
		chinese,
		func(ut ut.Translator) error {
			return ut.Add("sortby", "{0}字段必须是[{1}]中的一个", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("sortby", fe.Field(), fe.Param())
			return t
		},
	); nil != err {
		return
	}

	if err = validate.RegisterTranslation(
		"endswith_in",
		chinese,
		func(ut ut.Translator) error {
			return ut.Add("endswith_in", "{0}字段的后缀必须是[{1}]中的一个", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("endswith_in", fe.Field(), fe.Param())
			return t
		},
	); nil != err {
		return
	}

	return
}
