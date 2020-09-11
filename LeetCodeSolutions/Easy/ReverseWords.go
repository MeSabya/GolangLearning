package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {
	var ret []byte
	last := len(str) - 1
	for first := 0; first < last; first, last = first+1, last-1 {
		ret[first] = str[last]

	}
	fmt.Println(ret)
	return string(ret)
}
func reverseWords(s string) string {
	words := strings.Fields(s)
	fmt.Println("type of words %T", words)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	fmt.Println(" Logs 1")
	for i := 0; i < len(words); i++ {
		words[i] = reverse(words[i])
	}
	fmt.Println(" Logs 2")
	return strings.Join(words, " ")

}

func main() {
	fmt.Println(reverseWords("sabya sachi nayak"))
}
