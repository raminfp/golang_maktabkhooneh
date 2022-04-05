package main

import "fmt"

// global variable
var name string
var age int = 28

func main() {

	// explicit
	var everything string = "this is a test"
	// inferred
	var firstname, lastname = "ramin", "farajpour"
	// shorthand
	fname := "ramin"
	age := 20
	lname := "farajpour cami"
	// constant
	const address = "Iran, Tehran, St Test"
	fmt.Println(everything, name, age)
	fmt.Println(firstname, lastname)
	fmt.Println(fname, lname, age)
	fmt.Println(address)

}
