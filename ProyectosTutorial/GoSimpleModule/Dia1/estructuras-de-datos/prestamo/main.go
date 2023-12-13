package main

import "fmt"

func main() {
	var edad int = 35
	var empleado bool = true
	var antiguedad int = 4
	var sueldo int = 200000

	switch {
		case edad > 22, empleado, antiguedad > 1:
			fmt.Println("Puede acceder a un prestamo")
			fallthrough
		case sueldo > 100000:
			fmt.Println("No se cobrara interes")
		default:
			fmt.Println("No puede acceder a prestamo")
	}
}