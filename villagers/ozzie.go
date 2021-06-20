package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	ozzie string = "Ozzie"
)

var (
	// Ozzie is an Animal Crossing villager.
	//
	// Ozzie is a Koala.
	Ozzie Villager = villager{
		animal: animals.Koala.Name(),
		name:   ozzie}
)
