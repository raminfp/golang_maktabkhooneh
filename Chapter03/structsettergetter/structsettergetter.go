package main

import "fmt"

type Books struct {
	id     int
	Name   string
	Author string
}

//Getter

func (b Books) GetterBooks() Books {
	return b
}

// Setter

func (b *Books) SetterBooks(id int, author string, name string) {
	b.id = id
	b.Author = author
	b.Name = name
}

func main() {
	var book Books
	book.SetterBooks(1, "ali", "life")
	fmt.Println(book.GetterBooks().id, book.GetterBooks().Name, book.GetterBooks().Author)
	book.SetterBooks(2, "reza", "animals")
	fmt.Println(book.GetterBooks().id, book.GetterBooks().Name, book.GetterBooks().Author)
	book.SetterBooks(3, "mohsen", "dead")
	//fmt.Println(book.GetterBooks())
	fmt.Println(book.GetterBooks().id, book.GetterBooks().Name, book.GetterBooks().Author)

}
