package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Alligator is the species information for Alligators.
	Alligator Species = species{
		class:        reptilia,
		conservation: vulnerable,
		domain:       eukarya,
		family:       alligatoridae,
		fictional:    (!fictional),
		genus:        alligator,
		kingdom:      animalia,
		name:         animals.Alligator.Name(),
		order:        crocodylia,
		phylum:       chordata,
		species:      na}
)
