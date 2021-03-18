package main

import (
	"fmt"
)

type Coordinates struct {
	XCoord int
	YCoord int
}

func main() {

	sampleCoord1 := Coordinates{
		XCoord: 1,
		YCoord: -1,
	}

	sampleCoord2 := Coordinates{
		XCoord: -24,
		YCoord: 24,
	}

	sampleCoord3 := Coordinates{
		XCoord: 45,
		YCoord: 34,
	}

	sampleCoord4 := Coordinates{
		XCoord: 3,
		YCoord: 2,
	}

	suppliedSlice := []Coordinates{
		sampleCoord1, sampleCoord2, sampleCoord3, sampleCoord4,
	}

	resp, err := calculate(suppliedSlice)
	if err != nil {
		return
	}
	fmt.Printf("Max Distance: %.2f\nCoordinates: [%v, %v]\n", resp.Max.Value, resp.Max.Coordinates[0],
		resp.Max.Coordinates[1])
	fmt.Println("------------------------")
	fmt.Printf("Min Distance: %.2f\nCoordinates: [%v, %v]\n", resp.Min.Value, resp.Min.Coordinates[0],
		resp.Min.Coordinates[1])

}
