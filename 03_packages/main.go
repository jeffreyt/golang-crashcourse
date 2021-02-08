package main

import (
	"fmt"
	"math"
	"github.com/jeffreyt/golang-crashcourse/03_packages/strutil"
)

func main() {

	// using imported math package

	x := 2.7

	fmt.Println(math.Floor(x))
	fmt.Println(math.Ceil(x))
	fmt.Println(math.Sqrt(x))


	// and custom built strutil Reverse function

	fmt.Println(strutil.Reverse("hello backwards"))

}