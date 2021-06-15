package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	teddy string = "Teddy"
)

var (
	// Teddy is an Animal Crossing animal.
	//
	// Teddy is a Bear.
	Teddy Villager = villager{
		animal: animals.Bear.Name(),
		name:   teddy}
)
