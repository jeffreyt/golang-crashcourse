package main

import "fmt"

func main() {

	//Array
	ids := []int{33,43,23,52,52,33,25}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n",i,id)
	}

	// if not using index
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum: ", sum)

	//Maps
	ages := map[string]int{"Jeff": 45, "Uncle": 53, "Bob": 99}

	for k, v := range ages {
		fmt.Printf("%s is %d years old\n", k, v)
	}

	//only keys work too
	for k := range ages {
		fmt.Printf("Happy Birthday %s!\n", k)
	}

}