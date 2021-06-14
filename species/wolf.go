package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Wolf is the species information for Wolves.
	Wolf Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       canidae,
		fictional:    (!fictional),
		genus:        canis,
		kingdom:      animalia,
		name:         animals.Wolf.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      na}
)
