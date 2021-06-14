package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Camel is the species information for Camels.
	Camel Species = species{
		class:        mammalia,
		conservation: domesticated,
		domain:       eukarya,
		family:       camelidae,
		fictional:    (!fictional),
		genus:        camelus,
		kingdom:      animalia,
		name:         animals.Camel.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      na}
)
