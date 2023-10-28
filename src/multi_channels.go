package main

import "fmt"

func message(msg string, c chan string) {
	c <- msg
}

func main16() {
	c := make(chan string, 2)

	c <- "Hello"
	c <- "World"

	fmt.Println(len(c), cap(c))
	close(c) // close channel

	// iterate over channel
	for msg := range c {
		fmt.Println(msg)
	}

	// select
	c1 := make(chan string)
	c2 := make(chan string)

	go message("Mensaje1", c1)
	go message("Mensaje2", c2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("email1: ", msg1)
		case msg2 := <-c2:
			fmt.Println("email2: ", msg2)
		}
	}
}
