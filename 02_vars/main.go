package main

import "fmt"

func main() {

	// explicit typing
	var name string = "Jeff"
	var age int = 36

	fmt.Println(name, age)

	// or type can be inferred
	var name2 = "JT"
	var age2 = 69

	fmt.Println(name2, age2)
	fmt.Printf("name:%T, age:%T\n", name2, age2)

	// shorthand (cannot be used outside function)
	name3 := "JeffreyT"
	age3 := 12345
	fmt.Println(name3, age3)
	
	// or even shorter
	name4, age4 := "Jeffy-boy", 48.3
	fmt.Println(name4, age4)

	// consts
	const isCool = true
	fmt.Println(isCool)

	// cannot be reassigned
	// isCool = false
}