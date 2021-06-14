package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Squirrel is the species information for Squirrels.
	Squirrel Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       sciuridae,
		genus:        na,
		kingdom:      animalia,
		name:         animals.Squirrel.Name(),
		order:        rodentia,
		phylum:       chordata,
		species:      na}
)
