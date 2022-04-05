package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var x float64 = 10.5
	var rbyte byte = 'A'

	a := int(x)
	b := 11
	c := a + b
	// print c
	fmt.Println(c)
	// print byte
	fmt.Println(rbyte)
	//Print Size
	fmt.Printf("Size: %d\n", unsafe.Sizeof(c))
	//Print Type
	fmt.Printf("Type: %s\n", reflect.TypeOf(c))
	//Print Character
	fmt.Printf("Character: %c\n", rbyte)

}
