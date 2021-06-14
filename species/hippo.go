package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Hippo is the species information for Hippos.
	Hippo Species = species{
		class:        mammalia,
		conservation: vulnerable,
		domain:       eukarya,
		family:       hippopotamidea,
		fictional:    (!fictional),
		genus:        hippopotamus,
		kingdom:      animalia,
		name:         animals.Hippo.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      na}
)
