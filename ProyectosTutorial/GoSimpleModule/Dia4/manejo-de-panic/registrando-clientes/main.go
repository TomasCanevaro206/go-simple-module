package main

import (
	"fmt"
	"os"
	"errors"
)

type Cliente struct {
	File os.File
	Name string
	ID int
	PhoneNumber int
	Home string
}

var Clientes = []Cliente{}

func existeCliente(c Cliente) (bool, error){
	var existe bool = false
	var errExiste error = nil

	defer func() {
        err := recover()
        if err != nil {
            fmt.Println(err)
        }
    }()

	for _,cliente := range Clientes {
		if c == cliente {
			existe, errExiste = true, errors.New("El cliente ya existe")
			panic("Error: client already exists")
		}
	}
	return existe, errExiste
}

func clienteValido(c Cliente) (bool, error){
	if c.Name != "" && c.ID != 0 &&
		c.PhoneNumber != 0 && c.Home != "" {
			return true, nil
	} else {
		return false, errors.New("El cliente ingresado no es valido")
	}
}

func main(){

	custFile, err1 := os.Open("cliente1.txt")
	if err1 != nil {
		panic("The indicated file was not found or is damaged")
	}

	cliente1 := Cliente{*custFile, "martin", 4, 5876, "CABA"}

	existeCliente, err1 := existeCliente(cliente1)
	clienteValido, err2 := clienteValido(cliente1)

	if existeCliente {
		panic(err1)
	} else {
		Clientes = append(Clientes, cliente1)
	}

	if !clienteValido {
		panic(err2)
	}

	defer func(a *os.File) {
        err := recover()
        if err != nil {
            fmt.Println(err)
			fmt.Println("Several errors were detected at runtime")
        }
		a.Close()
    }(custFile)

	fmt.Println(Clientes)
	fmt.Println("End of execution")
}