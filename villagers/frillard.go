package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	frillard string = "Frillard"
)

var (
	// Frillard is an Animal Crossing villager.
	//
	// Frillard is a Frill-Necked Lizard.
	Frillard Villager = villager{
		animal: animals.FrillNeckedLizard.Name(),
		name:   frillard}
)
