package main

import "fmt"

func main() {
	c := 'a'
	c = c + 3
	fmt.Printf("%c", c)

	message := "shalom"
	for i := 0; i < 6; i++ {
		c := message[i]
		fmt.Printf("%c\n", c)
	}
}
