package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	monty string = "Monty"
)

var (
	// Monty is an Animal Crossing villager.
	//
	// Monty is a Monkey.
	Monty Villager = villager{
		animal: animals.Monkey.Name(),
		name:   monty}
)
