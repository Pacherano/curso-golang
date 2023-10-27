package main

import (
	"fmt"
	"strconv"
)

func main5() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		println("Es un uno")
	} else {
		println("No es un uno")
	}

	// and
	if valor1 == 1 && valor2 == 2 {
		println("Es un uno y un dos")
	}

	// or
	if valor1 == 0 || valor2 == 2 {
		println("Es un cero o un dos")
	}

	// convertir texto a entero
	value, err := strconv.Atoi("45")
	if err != nil {
		fmt.Printf("No pude convertir el número: %v\n", err)
	} else {
		fmt.Printf("El valor convertido es: %d\n", value)
	}

	// es numero par
	numero := 11
	if numero%2 == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}

	// switch
	modulo := 4 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// switch con condicion
	switch valor := 6 % 2; valor {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// switch sin condicion
	valor := 200
	switch {
	case valor > 100:
		fmt.Println("Es mayor a 100")
	case valor < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
}
