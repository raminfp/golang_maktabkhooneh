package main

import "fmt"

type Auth interface {
	CheckAuth()
}

type Admin struct {
	Username string
	Password string
}

type User struct {
	ID              uint64
	FullName, Email string
}

func (user User) CheckAuth() {
	fmt.Println("This is normal user")
}

func (admin Admin) CheckAuth() {
	fmt.Println("This is admin user")
}

func CommandAuth(auth Auth) {
	// user.CheckAuth()
	auth.CheckAuth()
}

func main() {

	//var auth Auth
	//user := User{ID: 1, Email: "ramin@gmail.com", FullName: "test"}
	//admin := Admin{Username: "admin", Password: "123456"}
	//user.CheckAuth()
	//admin.CheckAuth()
	user := User{ID: 1, Email: "ramin@gmail.com", FullName: "test"}
	admin := Admin{Username: "admin", Password: "123456"}
	CommandAuth(user)
	CommandAuth(admin)

}
