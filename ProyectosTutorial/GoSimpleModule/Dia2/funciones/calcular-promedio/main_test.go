package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPromedioNotas(t *testing.T) {

	notas := []int{2, 4, 6, 8}
	promedio := promedioNotas(notas...)
	var resultadoEsperado int = 5
	assert.Equal(t, resultadoEsperado, promedio, "El calculo del promedio fue incorrecto")
 }