package main

import (
	"fmt"
	"time"
)

func main() {
	const layot = "Mon, Jan 2, 2006"

	day := time.Now()
	tomorrow := day.Add(24 * time.Hour)

	fmt.Println(day.Format(layot))
	fmt.Println(tomorrow.Format(layot))
}
