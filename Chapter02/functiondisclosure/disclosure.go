package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func Add() func(string, int) {
	return func(s string, a int) {
		fmt.Println(s, a)
	}
}

func Add2() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// normal function
	hello()

	// anonymous function
	func(message string) {
		fmt.Println(message)
	}("Hello from anonymous function")

	// return anonymous function
	add := Add()
	add("Hello from return anonymous function", 12)

	add2 := Add2()
	fmt.Println(add2())
	fmt.Println(add2())
	fmt.Println(add2())
	fmt.Println(add2())

}
