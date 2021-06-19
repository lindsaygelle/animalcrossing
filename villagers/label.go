package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	label string = "Label"
)

var (
	// Label is an Animal Crossing villager.
	//
	// Label is a Hedgehog.
	Label Villager = villager{
		animal: animals.Hedgehog.Name(),
		name:   label}
)
