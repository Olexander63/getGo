package main

import "fmt"

func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haema", "Makemake", "Eris"}
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

}
