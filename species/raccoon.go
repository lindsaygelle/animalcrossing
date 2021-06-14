package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Raccoon is the species information for Raccoons.
	Raccoon Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       canidae,
		genus:        nyctereutes,
		kingdom:      animalia,
		name:         animals.Raccoon.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      nProcyonoides,
	}
)
