package iteration

import "strings"

func Repeat(char string, numRepeat int) (string, int) {
	var repeated strings.Builder
	var timesRepeated int

	for i := 0; i < numRepeat; i++ {
		repeated.WriteString(char)
		timesRepeated++
	}
	return repeated.String(), timesRepeated
}
