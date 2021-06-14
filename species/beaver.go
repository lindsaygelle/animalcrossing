package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Beaver is the species information for Beavers.
	Beaver Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       castoridae,
		fictional:    (!fictional),
		genus:        castor,
		kingdom:      animalia,
		name:         animals.Beaver.Name(),
		order:        rodentia,
		phylum:       chordata,
		species:      na}
)
