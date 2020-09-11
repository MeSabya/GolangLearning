package main

import "fmt"

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() {
	fmt.Println("Circle area computed")
}

func (s square) area() {
	fmt.Println("Square area computed")
}

type shape interface {
	area()
}

func info(s shape) {
	s.area()
}

func main() {

	c := circle{}
	s := square{}

	//c.area() //Try uncommenting this and below line
	//s.area()
	info(c) //Try commenting this line
	info(s) //Try commenting this line
}
