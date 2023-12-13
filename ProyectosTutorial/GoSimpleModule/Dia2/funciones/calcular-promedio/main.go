package main

import "fmt"

func promedioNotas( notas ...int ) int {

	suma := 0

	for _, nota := range notas {
		suma += nota
	}

	promedio := suma / len(notas)
	return promedio

}

func main() {
	var notas []int
	var nota int

	fmt.Println("Ingrese una nota (-1 para salir): ")
	fmt.Scanf("%d", &nota)
	if nota < 0 {
		fmt.Println("La nota debe ser positiva")
	} else {
		notas = append(notas,nota)
	}
	for true {
		fmt.Println("Ingrese una nota: ")
		fmt.Scanf("%d", &nota)
		if nota == -1 {
			break
		}
		if nota < 0 {
			fmt.Println("La nota debe ser positiva")
			continue
		}
		notas = append(notas,nota)
	}

	fmt.Println(notas)

	promedio := promedioNotas(notas...)
	fmt.Printf("El promedio de las notas es: %d\n", promedio)
}