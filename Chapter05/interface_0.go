package main

import "fmt"

func main() {
	// dynamic
	//var i interface{}
	var i any
	i = []int{1, 2, 4}
	fmt.Println(i)
	i = 123
	fmt.Println(i)
	i = map[string]int{"age": 20}
	fmt.Println(i)
	i = true
	fmt.Println(i)
	i = 23.13
	fmt.Println(i)
	i = "hello"
	fmt.Println(i)
	i = nil
	fmt.Println(i)

	// static
	var x interface{} = 123
	n, ok := x.(int)
	fmt.Println(ok)
	n = x.(int)
	fmt.Println(n)

}
