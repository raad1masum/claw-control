package main

import (
	"fmt"
)

func main() {
	if bangBang(10.0, 10.0) != 0 {
		fmt.Println(bangBang(10.0, 10.0))
	}
}

func bangBang(targetDistance float64, currentDistance float64) float64 {
	varError := targetDistance - currentDistance

	if varError > 0 {
		varError++
	} else if varError < 0 {
		varError--
	}
	
	return varError
}
