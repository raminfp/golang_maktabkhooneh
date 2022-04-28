package mod

import "fmt"

type User struct {
	ID              uint64
	FullName, Email string
}

func (user User) CheckAuth() {
	fmt.Println("This is normal user")
}
