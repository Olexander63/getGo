package main

import "fmt"

func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haema", "Makemake", "Eris"}
	for i := 0; i < len(dwarfs); i++ {
		dwarfs := dwarfs[i]
		fmt.Println(i, dwarfs)
	}
}
