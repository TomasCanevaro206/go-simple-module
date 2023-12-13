
package main

import (
	"fmt"
	"errors"
)

func dogFunc(cant float64) float64 {
	return cant * 10
}

func catFunc(cant float64) float64 {
	return cant * 5
}

func hamsterFunc(cant float64) float64 {
	return cant * 0.25
}

func tarantulaFunc(cant float64) float64 {
	return cant * 0.15
}

func animal(an string) (func(float64) float64, error) {

	var animalFunc func(float64) float64

	switch an {
		case "dog":
			animalFunc = dogFunc
		case "cat":
			animalFunc = catFunc
		case "hamster":
			animalFunc = hamsterFunc
		case "tarantula":
			animalFunc = tarantulaFunc
		default:
			return nil, errors.New("Operacion invalida")
	}
	return animalFunc, nil
}

func main() {

	const (
		dog = "dog"
		cat = "cat"
		hamster = "hamster"
		tarantula = "tarantula"
	 )
	 
	animalDog, msg := animal(dog)
	amountDog := animalDog(10)
	fmt.Println(amountDog)
	fmt.Println(msg)

	animalCat, msg := animal(cat)
	amountCat := animalCat(10)
	fmt.Println(amountCat)
	fmt.Println(msg)

	animalHamster, msg := animal(hamster)
	amountHamster := animalHamster(10)
	fmt.Println(amountHamster)
	fmt.Println(msg)

	animalTarantula, msg := animal(tarantula)
	amountTarantula := animalTarantula(10)
	fmt.Println(amountTarantula)
	fmt.Println(msg)
}