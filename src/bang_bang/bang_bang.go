package main

import (
	"fmt"
	"os"
)

func main() {
	var f1 float64
	var f2 float64

	fmt.Println("Enter Current Distance: ")
	_, err1 := fmt.Scanf("%f1", &f1)

	if err1 != nil {
		fmt.Println(" ") 
	}

	fmt.Println("Enter Target Distance: ")
	_, err2 := fmt.Scanf("%f2", &f2)

	if err2 != nil {
		fmt.Println(" ") 
	}

	bangBang(f2, f1)
}

func bangBang(targetDistance float64, currentDistance float64) {
	varError := targetDistance - currentDistance

	for {
		if varError > 0 {
			varError--
			fmt.Println("Error:", varError)
		} else if varError < 0 {
			varError++
			fmt.Println("Error:", varError)
		} else if varError == 0 {
			fmt.Print("Setpoint Achieved!\nExiting controller with ")
			os.Exit(1)
		}
	}
}
