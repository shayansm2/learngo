package main

import (
	"fmt"
	"math"
)

func MakeDistOrigin(oX, oY float64) func(float64, float64) float64 {
	return func(x, y float64) float64 { return math.Sqrt(math.Pow(x-oX, 2) + math.Pow(y-oY, 2)) }
}

func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fB := fA()
	fmt.Print(fB())
	fmt.Print(fB())
}
