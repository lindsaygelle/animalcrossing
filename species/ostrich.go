package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Ostrich is the species information for Ostriches.
	Ostrich Species = species{
		class:        aves,
		conservation: leastConcern,
		domain:       eukarya,
		family:       struthionidae,
		genus:        struthio,
		kingdom:      animalia,
		name:         animals.Ostrich.Name(),
		order:        struthioniformes,
		phylum:       chordata,
		species:      sCamelus}
)
