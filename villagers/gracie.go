package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gracie string = "Gracie"
)

var (
	// Gracie is an Animal Crossing villager.
	//
	// Gracie is a Giraffe.
	Gracie Villager = villager{
		animal: animals.Giraffe.Name(),
		name:   gracie}
)
