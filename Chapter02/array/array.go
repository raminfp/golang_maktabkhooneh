package main

import "fmt"

func main() {

	///       0      1       2
	///   ------------------------
	///   |       |      |       |
	///   ------------------------
	var number [3]int
	number[0] = 10
	number[1] = 20
	//fmt.Println(number)
	for i, v := range number {
		//fmt.Println(i, v)
		number[i] = v + 10
	}
	//fmt.Println(number)
	var name = []string{"ali", "reza", "mamad"}
	fmt.Println(name)
	fmt.Println(name[0])
	fmt.Println(name[1])
	for _, v1 := range name {
		fmt.Println(v1)
	}
	name = append(name, "zahra")
	fmt.Println(name)

}
