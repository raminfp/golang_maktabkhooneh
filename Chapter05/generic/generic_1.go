package main

import "fmt"

type GenericValue[T any] []T

func (g GenericValue[T]) Print() {
	for _, v := range g {
		fmt.Println(v)
	}
}

func MyPrint[T any](g GenericValue[T]) {
	for _, v := range g {
		fmt.Println(v)
	}
}

func main() {
	g := GenericValue[int]{1, 2, 3, 4}
	g.Print()
	i := GenericValue[string]{"one", "two", "three", "four"}
	i.Print()
	//MyPrint(g)
	MyPrint(i)

}
