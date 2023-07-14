package main

import "math"

func MakeDistOrigin(oX, oY float64) func(float64, float64) float64 {
	return func(x, y float64) float64 { return math.Sqrt(math.Pow(x-oX, 2) + math.Pow(y-oY, 2)) }
}
