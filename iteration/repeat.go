package iteration

func Repeat(charecter string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += charecter
	}
	return
}
