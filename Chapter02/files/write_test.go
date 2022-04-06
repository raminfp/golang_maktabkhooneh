package main

import "testing"

func TestWriteFile(t *testing.T) {
	myMessage := "this is a test"
	err := WriteFile(myMessage)
	if err != nil {
		t.Error(err.Error())
	}
}
