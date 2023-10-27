package main

import "fmt"

func main6() {

	// defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Cinco")
			continue
		}
		fmt.Println(i)
	}

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Cinco")
			break
		}
		fmt.Println(i)
	}
}
