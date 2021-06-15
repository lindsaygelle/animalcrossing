package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	dozer string = "Dozer"
)

var (
	// Dozer is an Animal Crossing animal.
	//
	// Dozer is a Bear.
	Dozer Villager = villager{
		animal: animals.Bear.Name(),
		name:   dozer}
)
