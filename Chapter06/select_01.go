package main

import "fmt"

func main() {

	firstname := make(chan string)
	lastname := make(chan string)
	quit := make(chan string)
	go send(firstname, lastname, quit)
	recive(firstname, lastname, quit)
}

func send(f, l, q chan<- string) {
	data := []string{"amin", "amini", "ali", "mosavi", "mohmamd", "hosseini"}
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			f <- data[i]
		} else {
			l <- data[i]
		}
	}
	q <- "finished"
}

func recive(f, l, q <-chan string) {
	for {
		select {
		case v := <-f:
			fmt.Println("firstname is : ", v)
		case v := <-l:
			fmt.Println("lastname is : ", v)
		case v := <-q:
			fmt.Println("my program is ", v)
			return
		}
	}
}
