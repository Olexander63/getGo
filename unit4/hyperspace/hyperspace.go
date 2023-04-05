package main

import (
	"fmt"
	"strings"
)

// hyperspace removes the space surrounding world

func hyperspace(world []string) {
	for i := range world {
		world[i] = strings.TrimSpace(world[i])
	}
}

func main() {

	planets := []string{" Venus", "Earth", " Mars"}
	hyperspace(planets)
	fmt.Println(strings.Join(planets, ""))
}
