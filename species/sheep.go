package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Sheep is the species information for Sheep.
	Sheep Species = species{
		class:        mammalia,
		conservation: domesticated,
		domain:       eukarya,
		family:       bovidae,
		genus:        ovis,
		kingdom:      animalia,
		name:         animals.Sheep.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      oAries}
)
