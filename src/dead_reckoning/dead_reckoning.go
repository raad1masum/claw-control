package main

import (
    "fmt"
)

func main() {
	var f1 float64
	var f2 float64

	fmt.Println("Enter Distance: ")
	_, err1 := fmt.Scanf("%f1", &f1)

	if err1 != nil {
		fmt.Println(" ") 
	}

	fmt.Println("Enter Velocity: ")
	_, err2 := fmt.Scanf("%f2", &f2)

	if err2 != nil {
		fmt.Println(" ") 
	}

	fmt.Println("Output:", deadReckoning(f1, f2), " seconds")

}

func deadReckoning(distance float64, velocity float64) float64 {
	return distance / velocity
}