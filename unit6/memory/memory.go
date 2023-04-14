package main

import "fmt"

func main() {
	answer := 42
	fmt.Print(&answer)

	address := &answer
	fmt.Println(*address)

}
