package main

import (
	"fmt"
	"time"
)

func Numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func Alphabet() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

func main() {
	go Numbers()
	go Alphabet()

	time.Sleep(2500 * time.Millisecond)
	//fmt.Println("terminate main")
}
