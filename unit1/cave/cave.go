package main

import "fmt"

func main() {
	var room = "cave"
	if room == "cave" {
		fmt.Println("You find yourself in a dimly li cavern.")
	} else if room == "entrance" {
		fmt.Println("There is a caver entrance here and a path to the east.")
	} else if room == "mountain" {
		fmt.Println("There is  a cliff here. A path lead west down the mpuntain.")
	} else {
		fmt.Println("Everything is white.")
	}
}
