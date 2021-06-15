package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	drShrunk string = "Dr. Shrunk"
)

var (
	// DrShrunk is an Animal Crossing Villager.
	//
	// DrShrunk is an Axolotl.
	DrShrunk Villager = villager{
		animal: animals.Axolotl.Name(),
		name:   drShrunk}
)
