package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Mouse is the species information for Mice.
	Mouse Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       muridae,
		genus:        mus,
		kingdom:      animalia,
		name:         animals.Mouse.Name(),
		order:        rodentia,
		phylum:       chordata,
		species:      na}
)
