package main

import (
	"fmt"
)

const OvenTime = 40

func main() {
	fmt.Println(RemainingOvenTime(30))
	fmt.Println(PreparationTime(2))
	fmt.Println(ElapsedTime(3, 20))
}

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	var remainingTime int
	remainingTime = OvenTime - actualMinutesInOven
	return remainingTime
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	var preparationTime int
	preparationTime = 2 * numberOfLayers
	return preparationTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	var elapsedTime int
	elapsedTime = PreparationTime(numberOfLayers) + actualMinutesInOven
	return elapsedTime
}
