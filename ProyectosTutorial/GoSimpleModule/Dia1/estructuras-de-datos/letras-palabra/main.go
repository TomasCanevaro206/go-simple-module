package main

import "fmt"

func main() {
	var palabra string = "miercoles"

	fmt.Println(len(palabra))

	for n := 0; n < len(palabra); n++ {
        fmt.Println(string(palabra[n]))
    }
}