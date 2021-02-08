package main

import "fmt"

func main() {
	
	a := 5
	b := &a

	fmt.Println(a,b)
	fmt.Printf("a: %T, b: %T\n", a, b)

	// get the value at the address
	fmt.Println(a,*b)

	// update value using pointer
	*b = 20
	fmt.Println(a)

}