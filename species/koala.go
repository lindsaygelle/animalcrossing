package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Koala is the species information for Koalas.
	Koala Species = species{
		class:        mammalia,
		conservation: vulnerable,
		domain:       eukarya,
		family:       phascolarctidae,
		fictional:    (!fictional),
		genus:        phascolarctos,
		kingdom:      animalia,
		name:         animals.Koala.Name(),
		order:        diprotodontia,
		phylum:       chordata,
		species:      pCinereus}
)
