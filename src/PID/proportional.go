package main

import (
	"fmt"
)

func main() {
	fmt.Println("PID FTW")
	proportional(10.0, 5.0, 1.0)
}

func proportional(setpoint float64, currentPosition float64, kP float64) {
	varError := setpoint - currentPosition
	output := 0.0

	if varError != 0 {
		output = currentPosition * kP
	}

	fmt.Println("Output:", output)
}
