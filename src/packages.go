package main

import (
	pk "curso-golang/src/mypackage"
	"fmt"
)

func main11() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage("Hola mundo")
}
