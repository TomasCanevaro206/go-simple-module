package main

import (
	"fmt"
	"os"
)

func leerAarchivo() {
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println(err)
        }
    }()

    _, err := os.Open("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

    fmt.Println("el archivo fue leido correctamente")
}

func main(){
 	leerAarchivo()
	fmt.Println("ejecucion finalizada")
}