package mod

func Start() {

	user := User{ID: 1, Email: "ramin@gmail.com", FullName: "test"}
	admin := Admin{Username: "admin", Password: "123456"}
	CommandAuth(user)
	CommandAuth(admin)

}
