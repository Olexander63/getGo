package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func measureTemperature(sample int, sensor func() kelvin) {
	for i := 0; i < sample; i++ {
		k := sensor()
		fmt.Printf("%vo K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151 + 150))
}
func main() {
	measureTemperature(3, fakeSensor)
}
