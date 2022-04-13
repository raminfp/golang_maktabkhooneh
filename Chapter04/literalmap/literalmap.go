package main

import "fmt"

func main() {

	m1 := map[string]any{}
	m1["one"] = 1
	m1["two"] = "hello"
	m1["three"] = 12.32
	fmt.Println(m1["one"])
	fmt.Println(m1["two"])
	fmt.Println(m1["three"])

	m2 := m1
	m2["four"] = 4
	fmt.Println(m2, len(m2))
	fmt.Println(m1, len(m1))

	delete(m1, "one")
	fmt.Println(m1)
	fmt.Println(m2)

}
