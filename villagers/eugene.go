package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	eugene string = "Eugene"
)

var (
	// Eugene is an Animal Crossing villager.
	//
	// Eugene is a Koala.
	Eugene Villager = villager{
		animal: animals.Koala.Name(),
		name:   eugene}
)
