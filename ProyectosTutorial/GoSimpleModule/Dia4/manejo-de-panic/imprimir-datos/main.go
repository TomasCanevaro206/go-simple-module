package main

import (
	"fmt"
	"os"
)

func leerAarchivo() {
    custFile, err1 := os.Open("customers.txt")
	data := make([]byte, 100)
	count, err2 := custFile.Read(data)
	if err1 != nil && err2 != nil {
		panic("The indicated file was not found or is damaged")
	} else {
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
	}

	defer func(a *os.File) {
        err := recover()
        if err != nil {
            fmt.Println(err)
        }
		a.Close()
    }(custFile)
}

func main(){
 	leerAarchivo()
	fmt.Println("ejecucion finalizada")
}