package main

import (
	"math"
)

type ReturnValues struct {
	Max Val
	Min Val
}

type Val struct {
	Value       float64
	Coordinates []Coordinates
}

func calculate(suppliedSlice []Coordinates) (ReturnValues, error) {
	var (
		currMaxCoord1 Coordinates
		currMaxCoord2 Coordinates
		currMinCoord1 Coordinates
		currMinCoord2 Coordinates
	)

	n := len(suppliedSlice)
	var maxDistance float64
	minDistance := math.MaxFloat64
	for i, slice := range suppliedSlice {
		for j := 0; j < n; j++ {
			if j != i {

				d := distance(slice, suppliedSlice[j])
				if d > maxDistance {
					currMaxCoord1 = Coordinates{
						XCoord: slice.XCoord,
						YCoord: slice.YCoord,
					}
					currMaxCoord2 = Coordinates{
						XCoord: suppliedSlice[j].XCoord,
						YCoord: suppliedSlice[j].YCoord,
					}
					maxDistance = d
				}
				if d < minDistance {
					currMinCoord1 = Coordinates{
						XCoord: slice.XCoord,
						YCoord: slice.YCoord,
					}
					currMinCoord2 = Coordinates{
						XCoord: suppliedSlice[j].XCoord,
						YCoord: suppliedSlice[j].YCoord,
					}
					minDistance = d
				}
			}
		}
	}

	return ReturnValues{
		Max: Val{
			Value:       maxDistance,
			Coordinates: []Coordinates{currMaxCoord1, currMaxCoord2},
		},
		Min: Val{
			Value:       minDistance,
			Coordinates: []Coordinates{currMinCoord1, currMinCoord2},
		},
	}, nil
}
