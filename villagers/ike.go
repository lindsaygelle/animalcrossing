package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ike string = "Ike"
)

var (
	// Ike is an Animal Crossing animal.
	//
	// Ike is a Bear.
	Ike Villager = villager{
		animal: animals.Bear.Name(),
		name:   ike}
)
