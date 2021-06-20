package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	porter string = "Porter"
)

var (
	// Porter is an Animal Crossing villager.
	//
	// Porter is a Monkey.
	Porter Villager = villager{
		animal: animals.Monkey.Name(),
		name:   porter}
)
