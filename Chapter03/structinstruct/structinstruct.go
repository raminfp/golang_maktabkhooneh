package main

import "fmt"

type A struct {
	AName string
	B
}

type B struct {
	BName string
	C
}

type C struct {
	CName string
}

func main() {
	a := A{
		B: B{
			C: C{
				CName: "this is a C struct\n",
			},
			BName: "this is struct B\n",
		},
		AName: "this is struct A\n",
	}

	fmt.Println(a.CName, a.AName, a.BName)
}
