package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Giraffe is the species information of Giraffes.
	Giraffe Species = species{
		class:        mammalia,
		conservation: vulnerable,
		domain:       eukarya,
		family:       giraffidae,
		genus:        giraffa,
		kingdom:      animalia,
		name:         animals.Giraffe.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      gCamelopardalis}
)
