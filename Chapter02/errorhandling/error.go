package main

import (
	"errors"
	"fmt"
	"net/mail"
)

func IsValidEmailAddr(email string) (string, bool) {

	addr, err := mail.ParseAddress(email)
	customeError := errors.New("bad email\n")
	if err != nil {
		return customeError.Error(), false
	}
	return addr.Address, true
}

func main() {
	isCheck, status := IsValidEmailAddr("tes)()(t@gmail.com")
	if status {
		fmt.Println(isCheck)
	} else {
		fmt.Println(isCheck)
	}

}
