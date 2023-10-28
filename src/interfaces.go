package main

import "fmt"

type figura interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figura) {
	fmt.Println(f.area())
}

func main13() {
	myCuadrado := cuadrado{base: 2}
	fmt.Println(myCuadrado.area())

	myRectangulo := rectangulo{base: 2, altura: 4}
	fmt.Println(myRectangulo.area())

	fmt.Println("--------------------")

	calcular(myCuadrado)
	calcular(myRectangulo)

	fmt.Println("--------------------")

	// lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)

}
