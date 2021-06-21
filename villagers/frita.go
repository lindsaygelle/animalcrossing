package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	frita string = "Frita"
)

var (
	// Frita is an Animal Crossing villager.
	//
	// Frita is a Sheep.
	Frita Villager = villager{
		animal: animals.Sheep.Name(),
		name:   frita}
)
