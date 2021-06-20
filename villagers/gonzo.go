package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	gonzo string = "Gonzo"
)

var (
	// Gonzo is an Animal Crossing villager.
	//
	// Gonzo is a Koala.
	Gonzo Villager = villager{
		animal: animals.Koala.Name(),
		name:   gonzo}
)
