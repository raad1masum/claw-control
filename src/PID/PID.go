package main

import (
	"fmt"
)

func main() {
	var f1 float64
	var f2 float64
	var f3 float64
	var f4 float64
	var f5 float64

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

	fmt.Println("Enter kP gain: ")
	_, err3 := fmt.Scanf("%f3", &f3)

	if err3 != nil {
		fmt.Println(" ") 
	}

	fmt.Println("Enter kI gain: ")
	_, err4 := fmt.Scanf("%f4", &f4)

	if err4 != nil {
		fmt.Println(" ") 
	}

	fmt.Println("Enter kD gain: ")
	_, err5 := fmt.Scanf("%f5", &f5)

	if err5 != nil {
		fmt.Println(" ") 
	}

	proportional := proportionalAdjust(f1, f2, f3)
	integral := integralAdjust(f1, f2, f4)
	derivative := derivativeAdjust(f1, f2, f5)

	output := proportional + integral + derivative

	fmt.Println("Ouput:", output)
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

func derivativeAdjust(setpoint float64, currentPosition, kD float64) float64 {
	varError := setpoint - currentPosition
	tmpEror := varError

	derivative := (varError - tmpEror) / 0.00001

	output := derivative * kD

	return output
}