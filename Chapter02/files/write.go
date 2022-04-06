package main

import (
	"errors"
	"os"
)

func WriteFile(message string) error {
	file, err := os.Create("test.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(message)
	if err != nil {
		return err
	}
	return errors.New("fake error")
}
