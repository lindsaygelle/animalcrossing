package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Hippopotamus is the species information for Hippos.
	Hippopotamus Species = species{
		class:        mammalia,
		conservation: vulnerable,
		domain:       eukarya,
		family:       hippopotamidea,
		fictional:    (!fictional),
		genus:        hippopotamus,
		kingdom:      animalia,
		name:         animals.Hippopotamus.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      na}
)
