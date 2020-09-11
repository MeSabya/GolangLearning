//Usage of "defer" keyword
//Its kind of with statemnt in python

package main

import "fmt"

func foo() {
	defer func() {
		fmt.Println("Lambda or anonymous function executed ")
	}() // This will get executed after below statement
	fmt.Println("Foo executed")
}
func main() {
	defer foo() //Here it will be called after below print stmt
	fmt.Println("Main function , with defer keyword usage")

}
