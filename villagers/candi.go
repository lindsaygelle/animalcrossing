package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	candi string = "Candi"
)

var (
	// Candi is an Animal Crossing villager.
	//
	// Candi is a Mouse.
	Candi Villager = villager{
		animal: animals.Mouse.Name(),
		name:   candi}
)
