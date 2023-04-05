package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(planets)
	fmt.Println(len(planets))
	fmt.Println(len(planets[3]))
	fmt.Println(planets[3], planets[2])
}
