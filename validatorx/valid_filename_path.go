package validatorx

import (
	"path/filepath"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// 不含特殊字符的有效名字 并且一个中文算1个字符
func checkValidFileNamePath(fl validator.FieldLevel) bool {
	return ValidFilenamePath(fl.Field().String())
}

// 有效文件名（Windows标准）/路径和文件名检测同一个 我们用200个字符
func ValidFilenamePath(file string) bool {
	if file == "/" {
		return true
	}
	fileRegexStr := `^[^\\\./:\*\?\"<>\|]{1}[^\\/:\*\?\"<>\|]{0,200}$`
	filenameRegex := regexp.MustCompile(fileRegexStr)
	f := filepath.Base(file)

	return filenameRegex.MatchString(f)
}

// 有效文件名（Windows标准）
func ValidFilename(filename string) bool {
	fileRegexStr := `^[^\\\./:\*\?\"<>\|]{1}[^\\/:\*\?\"<>\|]{0,254}$`
	filenameRegex := regexp.MustCompile(fileRegexStr)

	return filenameRegex.MatchString(filename)
}

// 有效文件夹名
func ValidFilepath(filepath string) bool {
	pathRegexStr := `^[^\\\/\?\*\&quot;\'\&gt;\&lt;\:\|]*$`
	pathRegex := regexp.MustCompile(pathRegexStr)
	if len(filepath) == 0 {
		return false
	}
	filepath = filepath[1:]

	return pathRegex.MatchString(filepath)
}
