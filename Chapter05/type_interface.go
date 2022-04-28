package main

import "fmt"

//declare interface
type MyInterface interface {
	// methods return
	SampleMethod()
	//SampleMethod() int
}

type Test struct{}

func (test Test) SampleMethod() {
	fmt.Println("This is a test")
}

func main() {
	var myIntface MyInterface

	//var tt Test
	//tt.SampleMethod()
	fmt.Println(myIntface)
	myIntface = Test{}
	myIntface.SampleMethod()

	//myIntface.SampleMethod()
}
