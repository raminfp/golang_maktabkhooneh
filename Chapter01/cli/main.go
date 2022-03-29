package main

//import (
//	"hello/db"
//	srv "hello/services"
//)
//
//func main() {
//	db.Connection()
//	srv.Email()
//}

// Direct import
//import "fmt"
//import "math"

// Group import
//import (
//	"fmt"
//	"math"
//)

// Nested import
//import (
//	"fmt"
//	"math/rand"
//)

//Alias Import
//import (
//	database 	"hello/db"
//	srv 		"hello/services"
//)
//
//func main() {
//	database.Connection()
//	srv.Email()
//}

// Ignore Import
import _ "fmt"
import "math"

func main() {
	math.Exp2(3)
}
