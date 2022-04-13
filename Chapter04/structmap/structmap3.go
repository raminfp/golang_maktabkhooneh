package main

import "fmt"

type Book struct {
	BookName  string
	TotalPage int
}

type Author struct {
	Name string
}

func main() {
	book := make(map[Author]Book)
	book[Author{Name: "Reza"}] = Book{
		BookName:  "Animal",
		TotalPage: 200,
	}
	book[Author{Name: "Amin"}] = Book{
		BookName:  "Dead",
		TotalPage: 300,
	}
	for k, _ := range book {
		fmt.Println("Book Name : ", book[k].BookName, " = Total Page : ", book[k].TotalPage)
	}
}
