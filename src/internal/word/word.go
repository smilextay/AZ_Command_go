package word

import (
	"strings"
	"unicode"
)

//ToUpper 转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

//ToLower 转换为小写
func ToLower(s string) string {

	return strings.ToLower(s)
}

//UnderScoreToUpperCameCase ..
func UnderScoreToUpperCameCase(s string) string {
	s = removeUnderScore(s)
	return strings.ToUpper(s)
}

//UnderScoreToLowerCameCase ..
func UnderScoreToLowerCameCase(s string) string {
	s = removeUnderScore(s)
	return strings.ToLower(s)
}

//CamelCaseToUnderScore 驼峰单词转为下划线
func CamelCaseToUnderScore(s string) string {

	var output = []rune{}
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

func removeUnderScore(s string) string {
	s = strings.Replace(s, "_", "", -1)
	return strings.Title(s)
}
