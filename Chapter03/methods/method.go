package main

import "fmt"

type Person struct {
	Name    string
	Address string
	Age     int
}

type Student struct {
	Firstname string
	Lastname  string
}

func (p *Person) ByName(input *Student) string {
	p.Name = "Ali"
	p.Address = "Iran - Tehran"
	p.Age = 30
	if input.Firstname == p.Name {
		return "is matched"
	}
	return "not match"
}

func (p *Person) ByChanageName() {
	p.Name = "Ali"
	p.Address = "Iran - Tehran"
	p.Age = 30
}

func main() {
	person := Person{}
	person.ByChanageName()
	fmt.Println(person)
	stud := &Student{Firstname: "ali"}
	fmt.Println(person.ByName(stud))

}
