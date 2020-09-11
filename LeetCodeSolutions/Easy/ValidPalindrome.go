package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)

	for _, rune := range s {
		n--
		runes[n] = rune
	}
	fmt.Println(string(runes))
	return string(runes[n:])

}
func main() {
	s := "madam"
	isalphanum := func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return -1
	}

	s2 := strings.ToLower(strings.Map(isalphanum, s))
	fmt.Println(s2 == reverse(s2))
}
