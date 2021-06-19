package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	winnie string = "Winnie"
)

var (
	// Winnie is an Animal Crossing villager.
	//
	// Winnie is a Horse.
	Winnie Villager = villager{
		animal: animals.Horse.Name(),
		name:   winnie}
)
