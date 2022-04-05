package main

import "testing"

func TestAdd(t *testing.T) {

	total := Add(4, 3)
	if total != 7 {
		t.Error("Add error")
	}
}

func TestSub(t *testing.T) {
	sub := sub(5, 10)
	if sub != 5 {
		t.Error("error")
	}
}
