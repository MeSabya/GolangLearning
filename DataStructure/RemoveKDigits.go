package main

import (
	"bytes"
	"fmt"
)

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	for idx := 0; idx < len(num); idx++ {
		for len(stack) > 0 && k > 0 && int(stack[len(stack)-1]) > int(num[idx]) {
			stack = stack[:len(stack)-1]
			k = k - 1
		}
		stack = append(stack, num[idx])
	}

	if k > 0 {
		for len(stack) > 0 && k > 0 {
			stack = stack[:len(stack)-1]
			k = k - 1
		}
	}

	idx := 0
	for idx < len(stack) && stack[idx] == '0' {
		idx = idx + 1
	}

	//stack = stack[idx:]
	fmt.Println("stack value is ", stack, idx)

	if len(stack[idx:]) > 0 {
		return bytes.NewBuffer(stack[idx:]).String()
	} else {
		return "0"
	}

}

func main() {
	fmt.Println(removeKdigits("10", 1))
}
