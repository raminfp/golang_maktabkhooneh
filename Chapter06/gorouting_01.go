package main

import (
	"fmt"
	"time"
)

//T 0 0.1 0.2 0.3 .....                     0.99
//  ---------------------------------------------------- ---
// 	17                                          22   19	 22
//  ---------------------------------------------------- ---

func sayHello(s string) {
	fmt.Println(s)
}

func main() {
	sayHello("A")
	sayHello("B")
	go sayHello("C")
	time.Sleep(1 * time.Millisecond)
	fmt.Println("terminate main")
}
