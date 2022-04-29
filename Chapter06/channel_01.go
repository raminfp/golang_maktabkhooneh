package main

import "fmt"

func main() {
	c := make(chan int, 10)
	go func() {
		//c <- 10 // write to channel
		for i := 1; i <= 10; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v) // read of channel
	}
}
