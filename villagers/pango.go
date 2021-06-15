package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	pango string = "Pango"
)

var (
	// Pango is an Animal Crossing villager.
	//
	// Pango is an Anteater.
	Pango Villager = villager{
		animal: animals.Anteater.Name(),
		name:   pango}
)
