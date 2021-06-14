package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Panther is the species information for Panthers.
	Panther Species = species{
		class:        mammalia,
		conservation: unknown,
		domain:       eukarya,
		family:       felidae,
		genus:        panthera,
		kingdom:      animalia,
		name:         animals.Panther.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      na}
)
