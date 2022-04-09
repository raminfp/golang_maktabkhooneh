package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
}

func main() {
	person := Person{
		Firstname: "A",
		Lastname:  "B",
	}
	fmt.Println(person.Firstname)
	person_2 := &Person{
		Firstname: "C",
		Lastname:  "D",
	}
	fmt.Println(person_2.Firstname)
	fmt.Println((*person_2).Firstname)
	//fmt.Println(*person)
	var a *Person = &person
	fmt.Println(a)

}
