package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	tex string = "Tex"
)

var (
	// Tex is an Animal Crossing villager.
	//
	// Tex is a Penguin.
	Tex Villager = villager{
		animal: animals.Penguin.Name(),
		name:   tex}
)
