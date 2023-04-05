package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Selector ZZ9",
		"MArs":  "Sector ZZ9",
	}
	planetsMArkII := planets
	planets["Earth"] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMArkII)

	delete(planets, "Earth")
	fmt.Println(planetsMArkII)
}
