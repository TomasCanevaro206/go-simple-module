package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestImpuestoSalario(t *testing.T) {

	var salario1 float64 = 75000
	impuesto1 := impuestoSalario(salario1)
	var resultadoEsperado1 float64 = 12750
	assert.Equal(t, resultadoEsperado1, impuesto1, "El calculo del impuesto fue incorrecto")

	var salario2 float64 = 200000
	impuesto2 := impuestoSalario(salario2)
	var resultadoEsperado2 float64 = 54000
	assert.Equal(t, resultadoEsperado2, impuesto2, "El calculo del impuesto fue incorrecto")

	var salario3 float64 = 40000
	impuesto3 := impuestoSalario(salario3)
	var resultadoEsperado3 float64 = 0
	assert.Equal(t, resultadoEsperado3, impuesto3, "El calculo del impuesto fue incorrecto")
 }
