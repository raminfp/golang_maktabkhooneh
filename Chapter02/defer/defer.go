package main

import "fmt"

// function
func myfunc(x int) {
	fmt.Println(x)
}

// method
//func (a struct) funx()  {
//
//}

func main() {

	// defer db.close()
	// connection database
	// while

	var a int = 10
	defer fmt.Println(a)
	a = 20
	fmt.Println(a)

	for i := 0; i < 10; i++ {
		defer myfunc(i)
	}

	b := 30
	fmt.Println(b)

	// db.close()
}
