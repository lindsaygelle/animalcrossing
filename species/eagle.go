package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Eagle is the species information for Eagles.
	Eagle Species = species{
		class:        aves,
		conservation: leastConcern,
		domain:       eukarya,
		family:       accipitridae,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Eagle.Name(),
		order:        accipitriformes,
		phylum:       chordata,
		species:      na}
)
