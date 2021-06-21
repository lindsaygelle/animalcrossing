package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// SeaGull is the species information for Sea Gulls.
	SeaGull Species = species{
		class:        aves,
		conservation: leastConcern,
		domain:       eukarya,
		family:       laridae,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.SeaGull.Name(),
		order:        charadriiformes,
		phylum:       chordata,
		species:      na}
)
