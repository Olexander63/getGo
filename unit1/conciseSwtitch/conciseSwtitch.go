package main

import "fmt"

func main() {
	fmt.Println("There isa cvaer entrance  here and a path to the east.")
	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You finde yourself i a dimly li cavern.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself i a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign read 'No Minors'.")
	default:
		fmt.Println("didn't quite get that.")
	}

}
