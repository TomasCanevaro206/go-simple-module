
package main

import (
	"fmt"
)

type Product struct {
	ID int
	Name string
	Price float64
	Description string
	Category string
}

var Products = []Product{
	{1, "Microondas", 200000, "un microondas para calentar comida", "cocina"},
	{2, "Aspiradora", 150000, "elemento de limpieza para casas", "limpieza"},
}

func (p Product) Save(productos []Product) {
	Products = append(productos, p)
	fmt.Println("Producto añadido")
}

func (p Product) GetAll(productos []Product) {
	for _,producto := range productos {
		fmt.Println(producto)
	}
}

func getById(id int) Product {
	var prod Product
	for _,producto := range Products {
		if producto.ID == id {
			prod = producto
		}
	}
	return prod
}

func main() {
	var producto3 = Product{3, "Afeitadora", 20000, "maquina para cortar pelo", "Higiene"}
	var producto4 = Product{4, "Reloj", 150000, "reloj para mirar hora y otras estadisticas", "Accesorios"}

	producto3.Save(Products)
	producto4.Save(Products)

	producto4.GetAll(Products)

	p1 := getById(1)
	fmt.Println(p1)

	p4 := getById(4)
	fmt.Println(p4)
}