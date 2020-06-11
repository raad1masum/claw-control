package main

import (
	"bufio"
    "fmt"
	"os"
	"strconv"
)

func main() {
	reader1 := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Distance: ")
	s1, _ := reader1.ReadString('\n')
	fmt.Println(s1)
	distance, _ := strconv.ParseFloat(s1, 64)
	fmt.Println(distance)

	reader2 := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Velocity: ")
	s2, _ := reader2.ReadString('\n')
	fmt.Println(s2)
	velocity, _ := strconv.ParseFloat(s2, 64)
	fmt.Println(velocity)

	fmt.Println("Output:", deadReckoning(distance, velocity))

}

func deadReckoning(distance float64, velocity float64) float64 {
	return distance / velocity
}