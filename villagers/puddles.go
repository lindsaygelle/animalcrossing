package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	puddles string = "Puddles"
)

var (
	// Puddles is an Animal Crossing villager.
	//
	// Puddles is a Frog.
	Puddles Villager = villager{
		animal: animals.Frog.Name(),
		name:   puddles}
)
