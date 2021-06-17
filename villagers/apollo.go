package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	apollo string = "Apollo"
)

var (
	// Apollo is an Animal Crossing villager.
	//
	// Apollo is an Eagle.
	Apollo Villager = villager{
		animal: animals.Eagle.Name(),
		name:   apollo}
)
