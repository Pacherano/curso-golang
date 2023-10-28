package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main12() {

	a := 2
	b := &a

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	fmt.Println("b:", *b)

	*b = 20
	fmt.Println("a:", a)

	fmt.Println("------------------")

	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRAM()
	fmt.Println(myPc)
	myPc.duplicateRAM()
	fmt.Println(myPc)

	fmt.Println("------------------")

	// formatear el string
	fmt.Println(myPc)
}
