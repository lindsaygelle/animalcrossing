package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Otter is the species information for Otters.
	Otter Species = species{
		class:        mammalia,
		conservation: nearThreatened,
		domain:       eukarya,
		family:       mustelidae,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Otter.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      na}
)
