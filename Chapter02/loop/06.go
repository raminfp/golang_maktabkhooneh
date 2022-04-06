package main

import "fmt"

func main() {
	/// while - infinity
	//for {
	//	time.Sleep(3 * time.Second)
	//	fmt.Println("while")
	//}

	/// for
	//var i int
	//for i = 0 ; i < 7; i++ {
	//	fmt.Println(i)
	//}

	//// while
	//i := 0
	//for i < 10 {
	//	i += 1
	//}
	//fmt.Println(i)

	/// for-each
	var name = [4]string{"golang", "python", "php", "java"}
	for i, v := range name {
		fmt.Println(i, v)
	}

	fmt.Println(len(name))
	fmt.Println(name[2])
	
}
