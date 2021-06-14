package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Pig is the species information for Pigs.
	Pig Species = species{
		class:        mammalia,
		conservation: domesticated,
		domain:       eukarya,
		family:       suidae,
		fictional:    (!fictional),
		genus:        sus,
		kingdom:      animalia,
		name:         animals.Pig.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      na}
)
