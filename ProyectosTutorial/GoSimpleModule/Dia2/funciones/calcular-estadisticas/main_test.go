package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcularSalario(t *testing.T) {

	var opMin string = "minimum"
	operacionMin, errMin := operation(opMin)
	valorMin := operacionMin(2, 3, 3, 4, 10, 2, 4, 5)
	var resultadoEsperadoMin int = 2
	assert.Equal(t, resultadoEsperadoMin, valorMin, "El calculo de la operacion fue incorrecto")
	assert.Nil(t, errMin)

	var opAve string = "average"
	operacionAve, errAve := operation(opAve)
	valorAve := operacionAve(2, 3, 3, 4, 1, 2, 4, 5)
	var resultadoEsperadoAve int = 3
	assert.Equal(t, resultadoEsperadoAve, valorAve, "El calculo de la operacion fue incorrecto")
	assert.Nil(t, errAve)

	var opMax string = "maximum"
	operacionMax, errMax := operation(opMax)
	valorMax := operacionMax(2, 3, 3, 4, 1, 2, 4, 5)
	var resultadoEsperadoMax int = 5
	assert.Equal(t, resultadoEsperadoMax, valorMax, "El calculo de la operacion fue incorrecto")
	assert.Nil(t, errMax)
 }