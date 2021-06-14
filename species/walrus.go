package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Walrus is the species information for Walruses.
	Walrus Species = species{
		class:        mammalia,
		conservation: unknown,
		domain:       eukarya,
		family:       odobenidae,
		genus:        odobenus,
		kingdom:      animalia,
		name:         animals.Walrus.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      oRosmarus}
)
