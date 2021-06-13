package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Kangaroo is the species information for Kangaroos.
	Kangaroo Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       macropodidae,
		genus:        macropus,
		kingdom:      animalia,
		name:         animals.Kangaroo.Name(),
		order:        diprotodontia,
		phylum:       chordata,
		species:      na}
)
