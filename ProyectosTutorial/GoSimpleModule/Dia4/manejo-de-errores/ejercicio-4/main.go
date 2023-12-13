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
		fmt.Println(fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %v", salary))
	}
}