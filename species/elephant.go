package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Elephant is the species information for Elephants.
	Elephant Species = species{
		class:        mammalia,
		conservation: endangered,
		domain:       eukarya,
		family:       elephantidae,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Elephant.Name(),
		order:        proboscidea,
		phylum:       chordata,
		species:      na}
)
