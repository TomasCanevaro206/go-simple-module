package main

import (
	"fmt"
)

type producto interface {
	Price() float64
}

type small struct {
	precio float64
}

func (s small) Price() float64{
	return s.precio
}

type medium struct {
	precio float64
}

func (m medium) Price() float64{
	return m.precio * 1.03
}

type large struct {
	precio float64
}

func (l large) Price() float64{
	return (l.precio * 1.06) + 2500
}

const (
	smallProd = "small"
	mediumProd = "medium"
	largeProd = "large"
)

func factory(tipoProd string, precio float64) producto {
	switch tipoProd {
	case smallProd:
		return small{precio}
	case mediumProd:
		return medium{precio}
	case largeProd:
		return large{precio}
	}
	return nil
}

func main() {
	s := factory("small", 100)
	fmt.Println("El precio del producto small es: ", s.Price())

	m := factory("medium", 100)
	fmt.Println("El precio del producto medium es: ", m.Price())

	l := factory("large", 100)
	fmt.Println("El precio del producto large es: ", l.Price())
}