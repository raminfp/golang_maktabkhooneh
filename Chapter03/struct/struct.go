package main

type Student struct {
	Firstname string
	Lastname  string
	Age       int
}

func MyStruct() (myStud Student) {
	student := Student{Firstname: "reza", Lastname: "mohamadi", Age: 10}
	//fmt.Println(student)
	student.Firstname = "ali"
	student.Lastname = "rezai"
	student.Age = 20
	//fmt.Println(student.Firstname, student.Lastname, student.Age)
	return student
}
