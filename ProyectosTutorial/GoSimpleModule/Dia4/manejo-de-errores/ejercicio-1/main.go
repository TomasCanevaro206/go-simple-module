package main

import (
	"fmt"
	"errors"
)

type errorSalario struct {}

func (e errorSalario) Error() error {
	return errors.New("Error: the salary entered does not reach the taxable minimum")
}

func tax(salary int) error {
	e := errorSalario{}
	if salary < 150000 {
		return e.Error()
	} else {
		fmt.Println("Must pay tax")
		return nil
	}
}

func main() {
	var salary int = 120000

	err := tax(salary)

	if(err != nil) {
		fmt.Println(err)
	}
}