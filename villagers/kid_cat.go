package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kidCat string = "Kid Cat"
)

var (
	// KidCat is an Animal Crossing villager.
	//
	// KidCat is a Cat.
	KidCat Villager = villager{
		animal: animals.Cat.Name(),
		name:   kidCat}
)
