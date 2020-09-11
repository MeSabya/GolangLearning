package main

import "fmt"

func test2() (int, string) {
	return 32, "sabya"
}
func test(fn func() (int, string)) {
	fmt.Println("test is called")
	age, name := fn()
	fmt.Println(age, name)

}

func funG(fn func(xi []int) int) int {
	arr := []int{1, 2, 3}
	return fn(arr)

}

func main() {
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return 1
	}
	test(test2)
	fmt.Println(funG(g))
}
