package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pedro"] = 20

	fmt.Println(m)

	//recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//obtener valor
	value := m["Jose"]
	fmt.Println(value)

	//obtener valor que no existe
	value2 := m["Jose2"]
	fmt.Println(value2)

	//obtener valor y si no existe
	value3, ok := m["Jose2"]
	fmt.Println(value3, ok)
	value4, ok := m["Jose"]
	fmt.Println(value4, ok)
}
