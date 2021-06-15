package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	klaus string = "Klaus"
)

var (
	// Klaus is an Animal Crossing animal.
	//
	// Klaus is a Bear.
	Klaus Villager = villager{
		animal: animals.Bear.Name(),
		name:   klaus}
)
