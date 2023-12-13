package main

import "fmt"

func calcularSalario( minutosTrabajados float64, categoria string ) float64 {

	var horasTrabajadas float64 = minutosTrabajados / 60
	var salario float64

	switch categoria {
		case "A":
			salario = (3000 * horasTrabajadas) * 1.5
		case "B":
			salario = (1500 * horasTrabajadas) * 1.2
		case "C":
			salario = 1000 * horasTrabajadas
		default:
			fmt.Println("categoria invalida")
	}

	return salario
}

func main() {

	var categoria string
	fmt.Println("Ingrese la categoria del empleado: ")
	fmt.Scanf("%s", &categoria)

	var minutosTrabajados float64
	fmt.Println("Ingrese los minutos trabajados en el mes: ")
	fmt.Scanf("%f", &minutosTrabajados)

	salario := calcularSalario(minutosTrabajados, categoria)
	fmt.Printf("Su salario sera: %f\n", salario)
}