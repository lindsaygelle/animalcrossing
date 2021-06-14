package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Penguin is the species information for Penguins.
	Penguin Species = species{
		class:        aves,
		conservation: endangered,
		domain:       eukarya,
		family:       spheniscidae,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Penguin.Name(),
		order:        sphenisciformes,
		phylum:       chordata,
		species:      na}
)
