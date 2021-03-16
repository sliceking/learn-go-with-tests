package iteration

import "strings"

// Repeat returns character repeated 5 times.
func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < 5; i++ {
		repeated.Write([]byte(character))
	}
	return repeated.String()
}
