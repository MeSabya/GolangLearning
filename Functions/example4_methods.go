//This is example of method

package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, p.age)
}
func main() {
	p := person{fname: "sabya", lname: "sachi", age: 31}
	p.speak()
}
