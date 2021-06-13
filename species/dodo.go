package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Dodo is the species information for Dodos.
	Dodo Species = species{
		class:        aves,
		conservation: extinct,
		domain:       eukarya,
		family:       columbidae,
		genus:        raphus,
		kingdom:      animalia,
		name:         animals.Dodo.Name(),
		order:        columbiformes,
		phylum:       chordata,
		species:      rCucullatus}
)
