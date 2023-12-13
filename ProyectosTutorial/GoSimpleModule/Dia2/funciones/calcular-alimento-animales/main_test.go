package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDog(t *testing.T) {

	var dog string = "dog"
	animalDog, msg := animal(dog)
	amount := animalDog(10)
	var resultadoEsperado float64 = 100
	assert.Equal(t, resultadoEsperado, amount, "El calculo de la cantidad de comida para perro fue incorrecto")
	assert.Nil(t, msg)
 }

 func TestCat(t *testing.T) {

	var cat string = "cat"
	animalCat, msg := animal(cat)
	amount := animalCat(10)
	var resultadoEsperado float64 = 50
	assert.Equal(t, resultadoEsperado, amount, "El calculo de la cantidad de comida para gato fue incorrecto")
	assert.Nil(t, msg)
 }

 func TestHamster(t *testing.T) {

	var hamster string = "hamster"
	animalHamster, msg := animal(hamster)
	amount := animalHamster(10)
	var resultadoEsperado float64 = 2.5
	assert.Equal(t, resultadoEsperado, amount, "El calculo de la cantidad de comida para hamster fue incorrecto")
	assert.Nil(t, msg)
 }

 func TestTarantula(t *testing.T) {

	var tarantula string = "tarantula"
	animalTarantula, msg := animal(tarantula)
	amount := animalTarantula(10)
	var resultadoEsperado float64 = 1.5
	assert.Equal(t, resultadoEsperado, amount, "El calculo de la cantidad de comida para tarantula fue incorrecto")
	assert.Nil(t, msg)
 }