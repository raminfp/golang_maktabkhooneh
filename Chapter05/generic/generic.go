package main

import "fmt"

func Print(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func genericPrint[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	Print([]string{"one", "two"})
	//Print([]int{1,2,3})
	genericPrint([]int{1, 2, 3, 4})
	genericPrint([]string{"1", "2"})

}
