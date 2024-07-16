package validatorx

import "testing"

func Test_ChineseOrEnglishOrNumbersOrSpace(t *testing.T) {
	t.Run(`valid_zh_space_num`, func(t *testing.T) {
		if !ChineseOrEnglishOrNumbersOrSpace(`测试 空格数字123`) {
			t.Errorf(`{测试 空格数字123}测试未通过`)
		}
	})

	t.Run(`valid_han_space_num_en`, func(t *testing.T) {
		if !ChineseOrEnglishOrNumbersOrSpace(`测试 空格数字123asasd`) {
			t.Errorf(`{测试 空格数字123asasd}测试未通过`)
		}
	})

	t.Run(`valid_han_num_en`, func(t *testing.T) {
		if !ChineseOrEnglishOrNumbersOrSpace(`撒大苏打123123efef`) {
			t.Errorf(`{撒大苏打123123efef}测试未通过`)
		}
	})

	t.Run(`valid_han_num`, func(t *testing.T) {
		if !ChineseOrEnglishOrNumbersOrSpace(`汉字123`) {
			t.Errorf(`{汉字123}测试未通过`)
		}
	})

	t.Run(`valid_han_space_num`, func(t *testing.T) {
		if !ChineseOrEnglishOrNumbersOrSpace(`汉字  123`) {
			t.Errorf(`{汉字  123}测试未通过`)
		}
	})

	t.Run(`invalid_han_num`, func(t *testing.T) {
		if ChineseOrEnglishOrNumbersOrSpace(`汉字123！@#`) {
			t.Errorf(`{汉字123！@#}测试未通过`)
		}
	})
}
