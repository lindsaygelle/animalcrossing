package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Hamster is the species information for Hamster.
	Hamster Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       cricetidae,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Hamster.Name(),
		order:        rodentia,
		phylum:       chordata,
		species:      na}
)
