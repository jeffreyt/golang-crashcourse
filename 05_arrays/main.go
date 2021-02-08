package main

import "fmt"

func main() {
	
	//Arrays should be typed and fixed length
	var fruit [2]string

	fruit[0] = "Apple"
	fruit[1] = "Orange"

	fmt.Println(fruit)

	//shorthand
	fruit2 := [2]string{"Melon", "Banana"}

	fmt.Println(fruit2)

	// doing this way you don't need to specify length
	fruit3 := []string{"Melon", "Banana", "Grape"}

	fmt.Println(fruit3)
	
	// useful stuff
	fmt.Println(len(fruit3))
	fmt.Println(fruit3[1:3])
}