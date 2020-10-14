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
	s1 := removeUnderScore(s)
	for i, _ := range s1 {
		s1[i] = strings.Title(s1[i])
	}
	s = strings.Join(s1, "")
	return strings.Replace(s, " ", "", -1)
}

//UnderScoreToLowerCameCase ..
func UnderScoreToLowerCameCase(s string) string {
	s = UnderScoreToUpperCameCase(s)
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

func removeUnderScore(s string) []string {
	return strings.Split(s, "_")
}
