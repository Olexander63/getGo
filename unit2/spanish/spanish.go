package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"
	fmt.Println(len(question), "bytes")
	// Prints 15 bytes
	fmt.Println(utf8.RuneCountInString(question), "runes")
	//Prints 12
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size)

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
}
