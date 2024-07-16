# validatorx
go-playground/validator的扩展，支持更多的检查，包括
- 手机号检查 mobile
- 密码检查 password
- 身份证检测 id_card,id_card_15
- 排序字段检查 sortby
- 前后字段空格检查 prefix_or_suffix_space
- 检查去掉'-数字'尾字符串后的字符串长度是否小于等于指定长度 并且一个中文算1个字符 max_len_without_number_suffix
- 只允许阳间名字 不允许含有^%&',;=?$ cn_en_num_space
- 字符串必须以字母开头 start_with_alpha