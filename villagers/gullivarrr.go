package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gullivarrr string = "Gullivarrr"
)

var (
	// Gullivarrr is an Animal Crossing villager.
	//
	// Gullivarrr is a Sea Gull.
	Gullivarrr Villager = villager{
		animal: animals.SeaGull.Name(),
		name:   gullivarrr}
)
