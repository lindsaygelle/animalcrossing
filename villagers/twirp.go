package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	twirp string = "Twirp"
)

var (
	// Twirp is an Animal Crossing villager.
	//
	// Twirp is a Bird.
	Twirp Villager = villager{
		animal: animals.Bird.Name(),
		name:   twirp}
)
