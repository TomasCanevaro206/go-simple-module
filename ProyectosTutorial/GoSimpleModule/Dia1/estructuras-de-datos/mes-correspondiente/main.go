package main

import "fmt"

func main() {
	meses := map[int]string{1: "enero", 2: "febrero", 3: "marzo", 4: "abril", 
	5: "mayo", 6: "junio", 7: "julio", 8: "agosto", 
	9: "septiembre", 10: "octubre", 11: "noviembre", 12: "diciembre"}

	var mes int
	fmt.Println("Ingrese un numero de mes: ")
	fmt.Scanln(&mes)
	mes = int(mes)

	if (mes < 1 || mes > 12) {
        fmt.Println("Por favor ingrese un mes valido")
    }

	fmt.Println("El mes ingresado es: ", meses[mes])
}