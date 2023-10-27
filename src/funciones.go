package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnFunction(a, b int) int {
	return a + b
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main3() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "hola")

	value := returnFunction(1, 2)
	fmt.Println(value)

	value1, value2 := doubleReturn(2)
	fmt.Println(value1, value2)

	value3, _ := doubleReturn(2)
	fmt.Println(value3)

}
