package main

import (
	"fmt"
	"runtime"
	"time"
)

func compare(x, y int) (z string) {
	if x > y {
		z = "x is greater than y"
	} else {
		z = "x is smaller than y"
	}
	return
}

func CheckOS() string {
	switch os := runtime.GOOS; os {
	case "linux":
		return os
	//case "windows":
	//	return os
	case "darwin":
		return os
	default:
		return os

	}
}

func TimeChecker() string {
	h := time.Now().Hour()
	switch {
	case h < 12:
		return "good morning"
	case h < 17:
		return "good afternoon"
	default:
		return "good night"
	}
}

func main() {
	t := compare(1, 3)
	fmt.Println(t)
	fmt.Println(CheckOS())
	fmt.Println(TimeChecker())
}
