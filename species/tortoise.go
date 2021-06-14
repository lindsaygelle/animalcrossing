package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Tortoise is the species information for Tortoises.
	Tortoise Species = species{
		class:        reptilia,
		conservation: unknown,
		domain:       eukarya,
		family:       testudinidae,
		genus:        na,
		kingdom:      animalia,
		name:         animals.Tortoise.Name(),
		order:        testudines,
		phylum:       chordata,
		species:      na,
	}
)
