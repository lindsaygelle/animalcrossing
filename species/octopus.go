package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Octopus is the species information for Octopuses.
	Octopus Species = species{
		class:        cephalopoda,
		conservation: leastConcern,
		domain:       eukarya,
		family:       na,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Octopus.Name(),
		order:        octopoda,
		phylum:       mollusca,
		species:      na}
)
