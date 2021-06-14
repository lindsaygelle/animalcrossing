package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Rabbit is the species information for Rabbits.
	Rabbit Species = species{
		class:        mammalia,
		conservation: domesticated,
		domain:       eukarya,
		family:       leporidae,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Rabbit.Name(),
		order:        lagomorpha,
		phylum:       chordata,
		species:      na,
	}
)
