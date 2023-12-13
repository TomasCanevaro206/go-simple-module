package main

import "fmt"

func impuestoSalario( salario float64 ) float64 {

	fmt.Println(salario)
	var impuesto float64

	if salario > 50000 && salario < 150000{
		impuesto = (salario * 17) / 100
		fmt.Println(impuesto)
	} else if salario > 150000 {
		impuesto = (salario * 27) / 100
	}else {
		impuesto = 0
	}

	return impuesto
}

func main() {
	var salario float64
	fmt.Println("Ingrese su salario: ")
	fmt.Scanln(&salario)
	impuesto := impuestoSalario(salario)
	fmt.Printf("El descuento a su salario sera: %f\n", impuesto)
}