package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Bear is the species information for Bears.
	Bear Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       ursidae,
		genus:        na,
		kingdom:      animalia,
		name:         animals.Bear.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      na}
)
