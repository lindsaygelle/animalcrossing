package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Deer is the species information for Deers.
	Deer Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       cervidae,
		genus:        na,
		kingdom:      animalia,
		name:         animals.Deer.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      na}
)
