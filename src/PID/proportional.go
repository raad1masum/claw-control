package main

import (
	"fmt"
)

func proportionalController() {
	var f1 float64
	var f2 float64
	var f3 float64

	fmt.Println("Enter Setpoint: ")
	_, err1 := fmt.Scanf("%f1", &f1)

	if err1 != nil {
		fmt.Println(" ") 
	}

	fmt.Println("Enter Target Position: ")
	_, err2 := fmt.Scanf("%f2", &f2)

	if err2 != nil {
		fmt.Println(" ") 
	}

	fmt.Println("Enter kP gain: ")
	_, err3 := fmt.Scanf("%f3", &f3)

	if err3 != nil {
		fmt.Println(" ") 
	}

	proportional(f1, f2, f3)
}

func proportional(setpoint float64, currentPosition float64, kP float64) {
	varError := setpoint - currentPosition
	output := 0.0

	if varError != 0 {
		output = currentPosition * kP
	}

	fmt.Println("Output:", output)
}
