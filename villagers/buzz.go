package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	buzz string = "Buzz"
)

var (
	// Buzz is an Animal Crossing villager.
	//
	// Buzz is an Eagle.
	Buzz Villager = villager{
		animal: animals.Eagle.Name(),
		name:   buzz}
)
