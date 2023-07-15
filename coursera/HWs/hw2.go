package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return (acceleration * math.Pow(time, 2) / 2) + (initialVelocity * time) + initialDisplacement
	}
}

func main() {
	var acceleration, initialVelocity, initialDisplacement float64
	fmt.Print("acceleration: ")
	fmt.Scanln(&acceleration)
	fmt.Printf("initial velocity: ")
	fmt.Scanln(&initialVelocity)
	fmt.Printf("initial displacement: ")
	fmt.Scanln(&initialDisplacement)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	var time float64
	fmt.Printf("time: ")
	fmt.Scanln(&time)
	fmt.Printf("current displacement: %v", fn(time))
}
