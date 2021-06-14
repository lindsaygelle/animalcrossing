package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Tiger is the species information for Tigers.
	Tiger Species = species{
		class:        mammalia,
		conservation: endangered,
		domain:       eukarya,
		family:       felidae,
		genus:        panthera,
		kingdom:      animalia,
		name:         animals.Tiger.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      pTigris}
)
