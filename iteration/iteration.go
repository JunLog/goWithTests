package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		// += 自增赋值运算符
		repeated += character
	}
	return repeated
}
