package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Pelican is the species information for Pelicans.
	Pelican Species = species{
		class:        aves,
		conservation: unknown,
		domain:       eukarya,
		family:       pelecanidae,
		fictional:    (!fictional),
		genus:        pelecanus,
		kingdom:      animalia,
		name:         animals.Pelican.Name(),
		order:        pelecaniformes,
		phylum:       chordata,
		species:      na}
)
