package main

import "fmt"

func main() {
	
	emails := make(map[string]string)

	// add entries
	emails["Jeff"] = "jeff@mail.com"
	emails["Uncle"] = "uncle@mail.com"	
	emails["Bob"] = "jimmy@mail.com"

	fmt.Println(emails)
	fmt.Println(emails["Uncle"])
	fmt.Println(len(emails))

	//delete
	delete(emails, "Bob")
	
	fmt.Println(emails)

}