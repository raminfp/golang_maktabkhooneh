package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyStruct(t *testing.T) {
	mytest := MyStruct()
	//if mytest.Age != 20 {
	//	t.Error("we have a error")
	//}
	//assert.Equal(t, 21, mytest.Age, "they should be equal")

	assert.NotEqual(t, 20, mytest.Age, "they should be equal")
}
