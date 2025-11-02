package main

import (
	"fmt"
)

func main() {
	var myName string = "Mehmet"
	myInt := 10
	myFloat := 10.0
	fmt.Printf("Hello %s, myInt is %d, myFloat is %f\n", myName, myInt, myFloat)

	// go has a concept of zero value - kinda default initial value?
	// string = "", bool = false, int = 0, float = 0
	// they can be represented in printf like below
	// string: %s
	// integer: %d
	// boolean: %t
	// float: %f
	var myInitialString string
	var myInitialInt int
	var myInitialFloat float32
	var myInitialBool bool
	fmt.Printf("myInitialString %v, myInitialInt is %v, myInitialFloat is %v, myInitialBool is %v\n", myInitialString, myInitialInt, myInitialFloat, myInitialBool)

}
