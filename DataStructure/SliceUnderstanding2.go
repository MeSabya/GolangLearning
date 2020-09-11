// This is most important program to understand slice internals
// https://medium.com/@riteeksrivastava/how-slices-internally-work-in-golang-a47fcb5d42ce
/*
REMEMBER:
When the slice length is equal to its capacity and you try to append a new element to it,
then the new memory is allocated
to the slice which has the double capacity as compared to the earlier capacity.
*/
package main

import "fmt"

func main() {
	a := []int{}
	a = append(a, 1)
	a = append(a, 2)
	b := append(a, 3)
	c := append(a, 4)

	fmt.Println("capacity and len of a", cap(a), len(a)) //Surprise surprise capacity and len of a 2 2
	fmt.Println("capacity and len of b", cap(b), len(b)) //capacity and len of b (4 3)
	fmt.Println("capacity and len of c", cap(c), len(c)) //capacity and len of c (4 3)

	a = append(a, 3)
	x := append(a, 4)
	y := append(a, 5)

	/************ understand here what is the vale of x and y , and why so ?*/
	fmt.Println("capacity and len of a", cap(a), len(a)) //Surprise surprise capacity and len of a 4 3
	fmt.Println("capacity and len of x", cap(x), len(x)) //capacity and len of x (4 4)
	fmt.Println("capacity and len of y", cap(y), len(y)) //capacity and len of y (4 4)

	z := append(x, 12)
	fmt.Println("*********************************************")
	fmt.Println("capacity and len of x", cap(x), len(x)) //capacity and len of x 4 4
	fmt.Println("capacity and len of z", cap(z), len(z)) //capacity and len of z 8 5
	fmt.Println("capacity and len of a", cap(a), len(a)) //capacity and len of a 4 3
	fmt.Println("Value of x", x)                         //Value of x [1 2 3 5]
	fmt.Println("Value of z", z)                         //Value of x [1 2 3 5 12]

}
