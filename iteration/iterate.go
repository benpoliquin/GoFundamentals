package iteration

import (
	"strings"
)

// Repeat the string the number of times repeatCount is set to
func Repeat(char string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += char
	}
	return repeated
}

// UseRepeatBuiltin using the builtin Repeat library
func UseRepeatBuiltin(char string, num int) string {
	return strings.Repeat(char, num)
}

func main() {
	UseRepeatBuiltin("a", 5)
}
