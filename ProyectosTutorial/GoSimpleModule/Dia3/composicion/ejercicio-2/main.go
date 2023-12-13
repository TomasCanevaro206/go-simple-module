
package main

import (
	"fmt"
)

type Person struct {
	ID int
	Name string
	DateOfBirth int
}

type Employee struct {
	Person
	ID int
	Position string
}

func (e Employee) PrintEmployee() {
	fmt.Println(e.ID)
	fmt.Println(e.Name)
	fmt.Println(e.DateOfBirth)
	fmt.Println(e.Position)
}

func main() {
	p := Person{14, "Tomas", 200695}
	e := Employee{p, 15, "Programmer"}

	e.PrintEmployee()
}