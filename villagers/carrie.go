package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	carrie string = "Carrie"
)

var (
	// Carrie is an Animal Crossing villager.
	//
	// Carrie is a Kangaroo.
	Carrie Villager = villager{
		animal: animals.Kangaroo.Name(),
		name:   carrie}
)
