package main

import (
	"fmt"
)

func main() {

	values := []interface{}{
		123, "123", true, 0.32, int32(876),
		map[int]bool{}, []int{1, 2, 3, 4, 5},
	}
	for _, x := range values {
		switch v := x.(type) {
		case []int:
			fmt.Println("int slice : ", v)
		case string:
			fmt.Println("string value : ", v)
		case bool:
			fmt.Println("bool value : ", v)
		case int:
			fmt.Println("int value : ", v)
		case int32:
			fmt.Println("int32 value : ", v)
		case nil:
			fmt.Println("nil value : ", v)
		default:
			fmt.Println("default value : ", v)
		}
	}

}
