package main

import "fmt"

func isPalindromo(text string) {
	var textReversed string

	for i := len(text) - 1; i >= 0; i-- {
		textReversed += string(text[i])
	}

	if text == textReversed {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

	// convert to lower case
	// if strings.ToLower(text) == strings.ToLower(textReversed) {
}

func main8() {
	isPalindromo("ana")
	isPalindromo("Ana")
}
