package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Peacock is the species information for Peacocks.
	Peacock Species = species{
		class:        aves,
		conservation: unknown,
		domain:       eukarya,
		family:       phasianidae,
		genus:        pavo,
		kingdom:      animalia,
		name:         animals.Peacock.Name(),
		order:        galliformes,
		phylum:       chordata,
		species:      na}
)
