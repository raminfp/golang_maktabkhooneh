package main

import "testing"

func TestReadFile(t *testing.T) {
	sample := "this is a test"
	data := ReadFile()
	if sample != data {
		t.Error(data)
	}
}
