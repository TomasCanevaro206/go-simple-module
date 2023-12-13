package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es : ", employees["Benjamin"])

	mayores := 0
	for emp, _ := range employees {
		if employees[emp] > 21 {
			mayores++
		}
	}
	fmt.Printf("Hay %d empleados mayores de 21 años\n", mayores)

	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println(employees)
}