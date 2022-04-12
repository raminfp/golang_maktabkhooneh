package main

import "fmt"

func main() {

	// mapp[key type][value type]
	//var m map[string]int
	m := make(map[string]int)
	m1 := make(map[string]string)
	m2 := make(map[string]float64)
	m1["one"] = "one"
	m2["one"] = 123.34
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	//m["one"] = 123
	fmt.Println(m["one"])
	fmt.Println(m1["one"])
	fmt.Println(m2["one"])

	count, ok := m["one"]
	if ok {
		fmt.Println("OK", count)
	} else if !ok {
		fmt.Println("Not ok ", count)
	}
	for key, value := range m {
		fmt.Println("Key : ", key, "Value : ", value)
	}

}
