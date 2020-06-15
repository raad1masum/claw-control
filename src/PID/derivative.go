package main

import (
	"fmt"
)

func main() {
	var f1 float64
	var f2 float64
	var f3 float64

	fmt.Println("Enter Setpoint: ")
	_, err1 := fmt.Scanf("%f1", &f1)

	if err1 != nil {
		fmt.Println(" ") 
	}

	fmt.Println("Enter Current Position: ")
	_, err2 := fmt.Scanf("%f2", &f2)

	if err2 != nil {
		fmt.Println(" ") 
	}

	fmt.Println("Enter kD gain: ")
	_, err3 := fmt.Scanf("%f3", &f3)

	if err3 != nil {
		fmt.Println(" ") 
	}

	derivative(f1, f2, f3)
}

func derivative(setpoint float64, currentPosition float64, kD  float64) {
	varError := setpoint - currentPosition
	tmpEror := varError

	derivative := (varError - tmpEror) / 0.00001

	output := derivative * kD
	fmt.Println("Output:", output)
}