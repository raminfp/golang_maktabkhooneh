package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return b - a
}

func swap(a, b int) (int, int) {
	return b, a
}

func Sum(values ...int64) (sum int64) {
	sum = 0
	for _, v := range values {
		sum += v
	}
	return
}

func Concat(str1 string, str2 string) string {
	return str1 + " " + str2
}

func main() {
	fmt.Println(Add(1, 3))
	//fmt.Println(sub(1, 3))
	//fmt.Println(swap(1, 3))
	a0 := Sum()
	a1 := Sum(2)
	a3 := Sum([]int64{2, 3, 5, 6}...)
	fmt.Println(a0, a1)
	fmt.Println(a3)
	fmt.Println(Concat("golang", "is awesome"))

}
