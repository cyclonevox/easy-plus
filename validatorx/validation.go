package validatorx

import (
	"github.com/go-playground/validator/v10"
)

func initValidation(validate *validator.Validate) (err error) {
	if err = validate.RegisterValidation("mobile", checkMobile); nil != err {
		return
	}
	if err = validate.RegisterValidation("password", checkPassword); nil != err {
		return
	}
	if err = validate.RegisterValidation("without_special_symbol", checkWithoutSpecialSymbol); nil != err {
		return
	}
	if err = validate.RegisterValidation("filename", checkValidFileNamePath); nil != err {
		return
	}
	if err = validate.RegisterValidation("start_with_alpha", checkStartWithAlpha); nil != err {
		return
	}
	if err = validate.RegisterValidation("prefix_or_suffix_space", checkPrefixOrSuffixSpace); nil != err {
		return
	}
	if err = validate.RegisterValidation("max_len_without_number_suffix", checkMaxStringLenWithoutNumberSuffix); nil != err {
		return
	}
	if err = validate.RegisterValidation("chinese", checkChinese); nil != err {
		return
	}
	if err = validate.RegisterValidation("name", checkName); nil != err {
		return
	}
	if err = validate.RegisterValidation("cn_en_num_space", chineseOrEnglishOrNumbersOrSpace); nil != err {
		return
	}
	if err = validate.RegisterValidation("sortby", checkSortBy); nil != err {
		return
	}
	if err = validate.RegisterValidation("id_card", checkIdCard); nil != err {
		return
	}
	if err = validate.RegisterValidation("id_card_15", checkIdCard15Len); nil != err {
		return
	}
	if err = validate.RegisterValidation("endswith_in", checkEndsWithIn); nil != err {
		return
	}

	return
}
