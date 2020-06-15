package main

import (
	"fmt"
)

func main() {
	derivative(10.0, 5.0, 2.0)
}

func derivative(setpoint float64, currentPosition float64, kD  float64) {
	varError := setpoint - currentPosition
	tmpEror := varError

	derivative := (varError - tmpEror) / 0.00001

	output := derivative * kD
	fmt.Println("Output:", output)
}