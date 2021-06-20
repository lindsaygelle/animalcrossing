package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	alice string = "Alice"
)

var (
	// Alice is an Animal Crossing villager.
	//
	// Alice is a Koala.
	Alice Villager = villager{
		animal: animals.Koala.Name(),
		name:   alice}
)
