package main

import (
	"math"
)

func distance(a, b Coordinates) float64 {
	first := math.Pow(float64(b.XCoord-a.XCoord), 2)
	second := math.Pow(float64(b.YCoord-a.YCoord), 2)

	return math.Sqrt(second + first)
}
