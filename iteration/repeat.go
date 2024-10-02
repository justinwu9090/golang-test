package iteration

func Repeat(character string) (repeated string) {
	// var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}
