package mod

import "fmt"

type Admin struct {
	Username string
	Password string
}

func (admin Admin) CheckAuth() {
	fmt.Println("This is admin user")
}
