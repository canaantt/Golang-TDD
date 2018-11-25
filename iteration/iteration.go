package iteration

func Repeat(str string, repeatCount int) string {
	result := ""
	for i := 0; i < repeatCount; i++ {
		result = result + str
	}
	return result
}
