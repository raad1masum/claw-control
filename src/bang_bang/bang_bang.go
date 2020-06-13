package main

import (
	"fmt"
	"os"
)

func main() {
	bangBang(10.0, 5.0)
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
			fmt.Print("Setpoint Achieved\nExiting controller with ")
			os.Exit(1)
		}
	}
}
