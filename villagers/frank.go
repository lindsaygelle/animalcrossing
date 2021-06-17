package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	frank string = "Frank"
)

var (
	// Frank is an Animal Crossing villager.
	//
	// Frank is an Eagle.
	Frank Villager = villager{
		animal: animals.Eagle.Name(),
		name:   frank}
)
