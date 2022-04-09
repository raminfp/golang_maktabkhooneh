package main

import "fmt"

type Author struct {
	id   int
	Name string
	Book Books
}

type Books struct {
	id   int
	Name string
}

func main() {

	authors := []*Author{
		&Author{
			id:   1,
			Name: "mohsen",
			Book: Books{
				id:   1,
				Name: "life"},
		},
		&Author{
			id:   2,
			Name: "ali",
			Book: struct {
				id   int
				Name string
			}{id: 2, Name: "dead"},
		},
		&Author{
			id:   3,
			Name: "reza",
			Book: struct {
				id   int
				Name string
			}{id: 3, Name: "animals"},
		},
	}

	for _, author := range authors {
		fmt.Println(author.id, author.Book.Name, author.Name)
	}
}
