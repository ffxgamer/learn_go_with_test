package iteration

import "strings"

// Repeat 把输入的字母重复5遍
func Repeat(character string, count int) string {
	return strings.Repeat(character, count)
}
