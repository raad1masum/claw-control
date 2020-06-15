package main

import (
	"fmt"
)

func main() {
	integral(10.0, 5.0, 2.0)
}

func integral(setpoint float64, currentPosition float64, kI float64) {
	varError := setpoint - currentPosition
	integral := varError * 0.00001

	output := integral * kI

	fmt.Println("Output:", output)
}