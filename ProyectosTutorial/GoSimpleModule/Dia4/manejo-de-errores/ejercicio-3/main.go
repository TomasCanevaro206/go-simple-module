package main

import (
	"fmt"
	"errors"
)

type errorSalario struct {}

func (e errorSalario) Error() string {
	return "Error: salary is less than 10000"
}

var e errorSalario = errorSalario{}

func errorSal() error {
	return errorSalario{}
}

func main() {
	var salary int = 90000

	var err errorSalario

	if salary <= 100000 {
		err = e
	}

	if errors.Is(err, errorSal()) {
		fmt.Println(errors.New("Error: salary is less than 10000"))
	}
}