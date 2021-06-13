package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Owl is the species information for Owls.
	Owl Species = species{
		class:        aves,
		conservation: unknown,
		domain:       eukarya,
		family:       na,
		genus:        na,
		kingdom:      animalia,
		name:         animals.Owl.Name(),
		order:        strigiformes,
		phylum:       chordata,
		species:      na}
)
