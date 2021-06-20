package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	lottie string = "Lottie"
)

var (
	// Lottie is an Animal Crossing villager.
	//
	// Lottie is an Otter.
	Lottie Villager = villager{
		animal: animals.Otter.Name(),
		name:   lottie}
)
