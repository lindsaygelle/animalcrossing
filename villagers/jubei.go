package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	jubei string = "Jūbei"
)

var (
	// Jubei is an Animal Crossing villager.
	//
	// Jubei is a Lion.
	Jubei Villager = villager{
		animal: animals.Lion.Name(),
		name:   jubei}
)
