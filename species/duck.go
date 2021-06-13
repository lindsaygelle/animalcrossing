package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Duck is the species information for Ducks.
	Duck Species = species{
		class:        aves,
		conservation: leastConcern,
		domain:       eukarya,
		family:       anatidae,
		genus:        na,
		kingdom:      enimalia,
		name:         animals.Duck.Name(),
		order:        anseriformes,
		phylum:       chordata,
		species:      na}
)
