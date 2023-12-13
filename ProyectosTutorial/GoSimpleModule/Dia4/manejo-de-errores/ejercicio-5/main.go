package main

import (
	"fmt"
	"errors"
)

func calculoSalarioMensual(horasTrabajadas float64, valorHora float64) (float64, error) {	
	if(horasTrabajadas < 80){
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}

	salario := horasTrabajadas * valorHora
	if salario >= 150000 {
		salario = salario*0.9
	}
	return salario, nil
}

func main() {
	salario1, err1 := calculoSalarioMensual(160, 2000)
	if err1 == nil {
		fmt.Println(salario1)
	} else {
		fmt.Println(err1)
	}

	salario2, err2 := calculoSalarioMensual(50, 2000)
	if err2 == nil {
		fmt.Println(salario2)
	} else {
		fmt.Println(err2)
	}
}