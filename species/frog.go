package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Frog is the species information for Frogs.
	Frog Species = species{
		class:        amphibia,
		conservation: vulnerable,
		domain:       eukarya,
		family:       na,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Frog.Name(),
		order:        anura,
		phylum:       chordata,
		species:      na}
)
