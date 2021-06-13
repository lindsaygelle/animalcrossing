package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Alpaca is the species information for Alpacas.
	Alpaca Species = species{
		class:        mammalia,
		conservation: domesticated,
		domain:       eukarya,
		family:       camelidae,
		genus:        vicugna,
		kingdom:      animalia,
		name:         animals.Alpaca.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      na}
)
