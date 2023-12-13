package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcularSalario(t *testing.T) {

	var minutosTrabajados1 float64 = 6000
	var categoria1 string = "A"
	salario1 := calcularSalario(minutosTrabajados1, categoria1)
	var resultadoEsperado1 float64 = 450000
	assert.Equal(t, resultadoEsperado1, salario1, "El calculo del salario fue incorrecto")

	var minutosTrabajados2 float64 = 6000
	var categoria2 string = "B"
	salario2 := calcularSalario(minutosTrabajados2, categoria2)
	var resultadoEsperado2 float64 = 180000
	assert.Equal(t, resultadoEsperado2, salario2, "El calculo del salario fue incorrecto")

	var minutosTrabajados3 float64 = 6000
	var categoria3 string = "C"
	salario3 := calcularSalario(minutosTrabajados3, categoria3)
	var resultadoEsperado3 float64 = 100000
	assert.Equal(t, resultadoEsperado3, salario3, "El calculo del salario fue incorrecto")
 }