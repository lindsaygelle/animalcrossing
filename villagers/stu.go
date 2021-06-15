package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	stu string = "Stu"
)

var (
	// Stu is an Animal Crossing villager.
	//
	// Stu is a Bull.
	Stu Villager = villager{
		animal: animals.Bull.Name(),
		name:   stu}
)
