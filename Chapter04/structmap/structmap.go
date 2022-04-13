package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Firstname string
	Lastname  string
	Address   []string
}

func main() {
	person := map[string]Person{}
	person["data"] = Person{Firstname: "ali", Lastname: "mahdavi", Address: []string{"Tehran", "Shiraz"}}
	fmt.Println(person["data"])
	personJson, err := json.Marshal(person)
	if err != nil {
		return
	}
	fmt.Println(string(personJson))
	/////// List of person
	lstPerson := make(map[string][]*Person)
	lstPerson["data"] = []*Person{
		{Firstname: "ramin", Lastname: "fp", Address: []string{"Tehran"}},
		{Firstname: "reza", Lastname: "f2p", Address: []string{"Tehran"}},
		{Firstname: "mehdi", Lastname: "f2p", Address: []string{"Tehran"}},
		{Firstname: "mohsen", Lastname: "f2p", Address: []string{"Tehran"}},
		{Firstname: "kazem", Lastname: "fp", Address: []string{"Tehran"}},
		{Firstname: "amin", Lastname: "fp123", Address: []string{"Tehran"}},
		{Firstname: "armin", Lastname: "fp12", Address: []string{"Tehran"}},
	}
	lstPersonJson, err := json.Marshal(lstPerson)
	if err != nil {
		return
	}
	fmt.Println(string(lstPersonJson))

}
