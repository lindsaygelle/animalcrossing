package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Lion is the species information for Lions.
	Lion Species = species{
		class:        ammalia,
		conservation: vulnerable,
		domain:       eukarya,
		family:       felidae,
		fictional:    (!fictional),
		genus:        panthera,
		kingdom:      animalia,
		name:         animals.Lion.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      pLeo}
)
