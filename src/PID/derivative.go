package main

import (
	"fmt"
)

func main() {
	fmt.Println("PID FTW")
}

func derivative(setpoint float64, currentPosition float64, kD  float64) {
	varError := setpoint - currentPosition
	tmpEror := varError

	derivative := (varError - tmpEror) / 0.00001

	output := derivative * kD
	fmt.Println(output)
}