package main

import "fmt"

func main() {
	message := "shalom"
	c := message[5]
	fmt.Printf("%c\n", c)

	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star)

	smile := ''
	fmt.Printf("%c %[1]v\n", smile)

	acute := 'é'
	fmt.Printf("%c %[1]v\n", acute)
}
