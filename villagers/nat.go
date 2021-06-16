package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	nat string = "Nat"
)

var (
	// Nat is an Animal Crossing villager.
	//
	// Nat is a Chameleon.
	Nat Villager = villager{
		animal: animals.Chameleon.Name(),
		name:   nat}
)
