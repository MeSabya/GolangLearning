package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p *Person) changeMe() {
	p.address = "ODISHA"
	p.age = 30
	p.name = "SABYASACHI NAYAK"
}

func (p Person) display() {
	fmt.Println(p.name, p.age, p.address)
}
func main() {
	p := Person{name: "sabya", address: "BANGALORE", age: 32}
	p.changeMe()
	p.display()

}
