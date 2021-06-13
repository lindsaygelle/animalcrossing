package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Hedgehog is the species information for Hedgehogs.
	Hedgehog Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       erinaceidae,
		genus:        na,
		kingdom:      animalia,
		name:         animals.Hedgehog.Name(),
		order:        eulipotyphla,
		phylum:       chordata,
		species:      na}
)
