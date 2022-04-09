package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
}

func ChangeValue(p *Person) {
	p.Firstname = "reza"
	p.Lastname = "hosseni"
}

func main() {
	person := Person{
		Firstname: "ali",
		Lastname:  "mohamadi",
	}
	ChangeValue(&person)
	fmt.Println(person.Firstname, person.Lastname)
}
