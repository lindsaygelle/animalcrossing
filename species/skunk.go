package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Skunk is the species information for Skunks.
	Skunk Species = species{
		class:        mammalia,
		conservation: unknown,
		domain:       eukarya,
		family:       mephitidae,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Skunk.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      na}
)
