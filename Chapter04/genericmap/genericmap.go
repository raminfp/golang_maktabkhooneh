package main

import "fmt"

func SumInt(m map[string]int) int {
	var sum int
	for _, v := range m {
		sum += v
	}
	return sum
}

func GenericSumInt[K comparable, V int](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum

}

func main() {

	m := make(map[string]int)
	//m := map[string]any{}
	m["1"] = 100
	m["2"] = 100
	//fmt.Println(SumInt(m))
	fmt.Println(GenericSumInt(m))

}
