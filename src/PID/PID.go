package main

import (
	"fmt"
)

func main() {
	fmt.Println("PID FTW")
}

func proportionalAdjust(setpoint float64, currentPosition float64, kP float64) float64 {
	varError := setpoint - currentPosition
	output := 0.0

	if varError != 0 {
		output = currentPosition * kP
	}

	return output
}

func integralAdjust(setpoint float64, currentPosition, kI float64) float64 {
	varError := setpoint - currentPosition
	integral := varError * 0.00001

	output := integral * kI

	return output
}