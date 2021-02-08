package main

import "fmt"


// args and return type
func greeting(name string) string {
	return "Hello " + name
}

func getSum(x, y int) int {
	return x + y
}


func main() {
	fmt.Println(greeting("Jeff"))


	fmt.Println(getSum(4, 5))
}