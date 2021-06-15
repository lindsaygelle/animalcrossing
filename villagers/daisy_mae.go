package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	daisyMae string = "Daisy Mae"
)

var (
	// DaisyMae is an Animal Crossing villager.
	//
	// DaisyMae is a Boar.
	DaisyMae Villager = villager{
		animal: animals.Boar.Name(),
		name:   daisyMae}
)
