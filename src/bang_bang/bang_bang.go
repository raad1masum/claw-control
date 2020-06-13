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
			fmt.Println(varError)
		} else if varError < 0 {
			varError++
			fmt.Println(varError)
		} else if varError == 0 {
			fmt.Print("Setpoint achieved, exiting controller: ")
			os.Exit(1)
		}
	}
}
