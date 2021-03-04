package word

import (
	"strings"
	"unicode"
)

// ToUpper 全部转化为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 全部转化为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscopeToUpperCamelCase 下划线单词转大写驼峰
func UnderscopeToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下划线单词转小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscopeToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰单词转下划线单词
func CamelCaseToUnderscore(s string) string {
	var out []rune
	for i, r := range s {
		if i == 0 {
			out = append(out, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(r))
	}
	return string(out)
}
