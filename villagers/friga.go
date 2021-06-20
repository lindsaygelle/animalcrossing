package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	friga string = "Friga"
)

var (
	// Friga is an Animal Crossing villager.
	//
	// Friga is a Penguin.
	Friga Villager = villager{
		animal: animals.Penguin.Name(),
		name:   friga}
)
