package main

import "fmt"

type Person struct {
	Fisrtname string
	Lastname  string
}

func ChangeName(p Person) {
	p.Fisrtname = "mehdi"
	p.Lastname = "mosavi"
	//fmt.Println(p)
}

func main() {

	person := Person{
		Fisrtname: "ali",
		Lastname:  "mohamadi",
	}
	//fmt.Println(person)
	ChangeName(person)
	fmt.Println(person)

}
