package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Chameleon is the species information for Chameleons.
	Chameleon Species = species{
		class:        reptilia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       chamaeleonidae,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Chameleon.Name(),
		order:        squamata,
		phylum:       chordata,
		species:      na}
)
