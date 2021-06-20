package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	mathilda string = "Mathilda"
)

var (
	// Mathilda is an Animal Crossing villager.
	//
	// Mathilda is a Kangaroo.
	Mathilda Villager = villager{
		animal: animals.Kangaroo.Name(),
		name:   mathilda}
)
