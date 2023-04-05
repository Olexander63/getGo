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

	terrestrial := planets[0:4]
	gasGiant := planets[4:6]
	iceGiants := planets[6:8]
	fmt.Println(terrestrial, gasGiant, iceGiants)

	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(giants, gas, ice)

	iceGiantMarkII := iceGiants
	iceGiants[1] = "Poseidon"
	fmt.Println(planets)
	fmt.Println(iceGiants, iceGiantMarkII, ice)
}
