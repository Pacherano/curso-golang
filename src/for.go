package main

import "fmt"

func main4() {

	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------------")

	// For while
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	fmt.Println("------------------")

	// For forever
	// k := 0
	// for {
	// 	fmt.Println(k)
	// 	k++
	// }

	fmt.Println("------------------")

	// for inverso
	for l := 10; l > 0; l-- {
		fmt.Println(l)
	}
}
