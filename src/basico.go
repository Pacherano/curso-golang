package main

import "fmt"

func main2() {
	//hola mundo
	fmt.Println("Hola Mundo")

	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	//División
	result = y / x
	fmt.Println("División:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)

	//Retos
	//Area de un rectangulo
	baseRectangulo := 20
	alturaRectangulo := 10
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("Area rectangulo:", areaRectangulo)

	//Area de un trapecio
	baseMayorTrapecio := 20
	baseMenorTrapecio := 10
	alturaTrapecio := 10
	areaTrapecio := ((baseMayorTrapecio + baseMenorTrapecio) * alturaTrapecio) / 2
	fmt.Println("Area trapecio:", areaTrapecio)

	//Area de un circulo
	const piCirculo float64 = 3.14
	radioCirculo := 10
	areaCirculo := piCirculo * float64(radioCirculo*radioCirculo)
	fmt.Println("Area circulo:", areaCirculo)

	//Area de una elipse
	const piElipse float64 = 3.14
	radioMayorElipse := 10
	radioMenorElipse := 5
	areaElipse := piElipse * float64(radioMayorElipse*radioMenorElipse)
	fmt.Println("Area elipse:", areaElipse)

	//Area de un poligono
	perimetroPoligono := 20
	apotemaPoligono := 10
	areaPoligono := (perimetroPoligono * apotemaPoligono) / 2
	fmt.Println("Area poligono:", areaPoligono)

	//Area de una esfera
	radioEsfera := 10
	areaEsfera := 4 * piCirculo * float64(radioEsfera*radioEsfera)
	fmt.Println("Area esfera:", areaEsfera)

	//Area de un cilindro
	alturaCilindro := 10
	areaCilindro := 2 * piCirculo * float64(radioEsfera*alturaCilindro)
	fmt.Println("Area cilindro:", areaCilindro)

	//Area de un cono
	alturaCono := 10
	areaCono := piCirculo * float64(radioEsfera*(radioEsfera+alturaCono))
	fmt.Println("Area cono:", areaCono)

	//Area de un cubo
	areaCubo := 6 * float64(baseCuadrado*baseCuadrado)
	fmt.Println("Area cubo:", areaCubo)

	//Area de una piramide
	areaPiramide := float64(baseCuadrado*baseCuadrado) + float64((4*baseCuadrado)*apotemaPoligono)
	fmt.Println("Area piramide:", areaPiramide)

	//paquete fmt

	// declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
