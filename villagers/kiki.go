package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kiki string = "Kiki"
)

var (
	// Kiki is an Animal Crossing villager.
	//
	// Kiki is a Cat.
	Kiki Villager = villager{
		animal: animals.Cat.Name(),
		name:   kiki}
)
