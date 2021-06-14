package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Anteater is the species information for Anteaters.
	Anteater Species = species{
		class:        mammalia,
		conservation: vulnerable,
		domain:       eukarya,
		family:       na,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Anteater.Name(),
		order:        pilosa,
		phylum:       chordata,
		species:      na}
)
