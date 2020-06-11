package main

import (
	"fmt"
)

func main() {
	distance := 10.0
	velocity := 5.0
	
	fmt.Println("Output:", deadReckoning(distance, velocity))
}

func deadReckoning(distance float64, velocity float64) float64 {
	return distance / velocity
}