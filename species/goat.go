package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Goat is the species information for Goats.
	Goat Species = species{
		class:        mammalia,
		conservation: domesticated,
		domain:       eukarya,
		family:       bovidae,
		genus:        capra,
		kingdom:      animalia,
		name:         animals.Goat.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      cAegagrus}
)
