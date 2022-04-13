package main

import "fmt"

func main() {

	// array
	var myarray [2]string
	myarray[0] = "1"
	myarray[1] = "2"
	// slice
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sliceprint(a[:7])
	sliceprint(a[:0])

	animal := []string{"dog", "cat"}
	animal[0] = "test"
	fmt.Println(animal)

	b := make([]int, 10)
	b[0] = 0
	b[1] = 1
	b[2] = 2
	fmt.Println(b)

}

func sliceprint(param []int) {
	fmt.Println(param)
}
