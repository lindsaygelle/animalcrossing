package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	olaf string = "Olaf"
)

var (
	// Olaf is an Animal Crossing villager.
	//
	// Olaf is an Anteater.
	Olaf Villager = villager{
		animal: animals.Anteater.Name(),
		name:   olaf}
)
