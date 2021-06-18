package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	hamphrey string = "Hamphrey"
)

var (
	// Hamphrey is an Animal Crossing villager.
	//
	// Hamphrey is a Hamster.
	Hamphrey Villager = villager{
		animal: animals.Hamster.Name(),
		name:   hamphrey}
)
