
package main

import (
	"fmt"
)

type Alumno struct{
	Name string
	Apellido string
	DNI int
	Fecha int
}

func (a Alumno) detalle() {
	fmt.Println("Nombre: ", a.Name)
	fmt.Println("Apellido: ", a.Apellido)
	fmt.Println("DNI: ", a.DNI)
	fmt.Println("Fecha de ingreso: ", a.Fecha)
}

func main() {
	alumno := Alumno{"Martin", "Rosetti", 25467896, 12052021}
	alumno.detalle()
}