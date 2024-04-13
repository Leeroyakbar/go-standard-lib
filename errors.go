package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation Error")
	NotFoundError   = errors.New("Not Found Error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	} else if id != "Lee" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("Lee")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Printf("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Printf("not found error")
		} else {
			fmt.Println("Uknown Error")
		}
	}
}
