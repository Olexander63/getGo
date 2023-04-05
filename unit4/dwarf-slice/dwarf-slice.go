package main

import "fmt"

func main() {
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]

	fmt.Println(dwarfSlice)

	fmt.Printf("array %T\n", dwarfArray)
}
