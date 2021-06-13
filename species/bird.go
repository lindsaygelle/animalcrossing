package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Bird is the species information for Birds.
	Bird Species = species{
		class:        aves,
		conservation: leastConcern,
		domain:       eukarya,
		family:       na,
		genus:        na,
		kingdom:      animalia,
		name:         animals.Bird.Name(),
		order:        na,
		phylum:       chordata,
		species:      na}
)
