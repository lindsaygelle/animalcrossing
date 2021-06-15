package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kabuki string = "Kabuki"
)

var (
	// Kabuki is an Animal Crossing villager.
	//
	// Kabuki is a Cat.
	Kabuki Villager = villager{
		animal: animals.Cat.Name(),
		name:   kabuki}
)
