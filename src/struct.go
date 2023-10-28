package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main10() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// otra forma de declarar un struct
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
