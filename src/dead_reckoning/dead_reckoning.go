package main

import (
	"fmt"
)

func main() {
	fmt.Println("Dead Reckoning")
}

func deadReckoning(distance float64, velocity float64) float64 {
	return distance / velocity
}