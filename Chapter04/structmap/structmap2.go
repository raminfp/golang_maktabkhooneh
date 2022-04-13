package main

import (
	"encoding/json"
	"fmt"
)

type Person2 struct {
	Firstname string
	Lastname  string
	Address   []string
}

func main() {

	// data = {status:OK, person:{Name, address}}
	person := map[string]any{}
	person2 := map[string]any{}
	person["user"] = Person2{Firstname: "ali", Lastname: "mahdavi", Address: []string{"Tehran", "Shiraz"}}
	person["statusCode"] = 404
	personJson, err := json.Marshal(person)
	if err != nil {
		return
	}
	fmt.Println(string(personJson))

	lstPerson := make(map[string][]*Person2)
	lstPerson["data"] = []*Person2{
		{Firstname: "ramin", Lastname: "fp", Address: []string{"Tehran"}},
		{Firstname: "reza", Lastname: "f2p", Address: []string{"Tehran"}},
		{Firstname: "mehdi", Lastname: "f2p", Address: []string{"Tehran"}},
		{Firstname: "mohsen", Lastname: "f2p", Address: []string{"Tehran"}},
		{Firstname: "kazem", Lastname: "fp", Address: []string{"Tehran"}},
		{Firstname: "amin", Lastname: "fp123", Address: []string{"Tehran"}},
		{Firstname: "armin", Lastname: "fp12", Address: []string{"Tehran"}},
	}
	// map == any   =    map == person
	person2["info"] = lstPerson["data"]
	person2["statusCode"] = 200
	lstPersonJson, err := json.Marshal(person2)
	if err != nil {
		return
	}
	fmt.Println(string(lstPersonJson))
}
