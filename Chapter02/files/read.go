package main

import "io/ioutil"

func ReadFile() string {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return err.Error()
	}
	return string(data)

}
