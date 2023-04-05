package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Memake", "Eris"}
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)
}
