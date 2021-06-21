package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	zipperTBunny string = "Zipper T. Bunny"
)

var (
	// ZipperTBunny is an Animal Crossing villager.
	//
	// ZipperTBunny is a Rabbit.
	ZipperTBunny Villager = villager{
		animal: animals.Rabbit.Name(),
		name:   zipperTBunny}
)
