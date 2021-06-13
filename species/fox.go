package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Fox is the species information for Foxes.
	Fox Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       canidae,
		genus:        vulpes,
		kingdom:      animalia,
		name:         animals.Fox.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      vVulpes}
)
