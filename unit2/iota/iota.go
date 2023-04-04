package main

import (
	"fmt"
	"strconv"
)

func main() {
	countdwon := 9
	str := fmt.Sprintf("Launch int T minus %v seconds", countdwon)
	fmt.Println(str)

	countdown, err := strconv.Atoi("10")
	if err != nil {
		// oh no, something went wrong
	}
	fmt.Println(countdown)
}
