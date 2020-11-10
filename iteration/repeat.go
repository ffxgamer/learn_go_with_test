package iteration

// Repeat 把输入的字母重复5遍
func Repeat(character string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += character
	}
	return result
}
