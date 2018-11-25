package iteration

const repeatCount = 5

func Repeat(str string) string {
	result := ""
	for i := 0; i < repeatCount; i++ {
		result = result + str
	}
	return result
}
