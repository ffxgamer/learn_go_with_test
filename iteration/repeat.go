package iteration

// Repeat 把输入的字母重复5遍
func Repeat(character string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += character
	}
	return result
}
