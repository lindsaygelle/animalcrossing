package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	wLink string = "W. Link"
)

var (
	// WLink is an Animal Crossing villager.
	//
	// WLink is a Wolf.
	WLink Villager = villager{
		animal: animals.Wolf.Name(),
		name:   wLink}
)
