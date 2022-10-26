package main

import (
	"fmt"
)

func main() {
	var (
		knightIsAwake   bool
		archerIsAwake   bool
		prisonerIsAwake bool
		dogIsPresent    bool
	)
	knightIsAwake = false
	archerIsAwake = true
	prisonerIsAwake = true
	dogIsPresent = false
	fmt.Println(CanFastAttack(knightIsAwake))
	fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
	fmt.Println(CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
	fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, dogIsPresent))
}

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	IfDogIsThere := petDogIsPresent && !archerIsAwake
	IfDogIsNotThere := !petDogIsPresent && !knightIsAwake && !archerIsAwake && prisonerIsAwake
	return IfDogIsThere || IfDogIsNotThere
}
