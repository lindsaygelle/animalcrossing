package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Boar is the species information for Boars.
	Boar Species = species{
		class:        mammalia,
		conservation: leastConcern,
		domain:       eukarya,
		family:       suidae,
		genus:        sus,
		kingdom:      animalia,
		name:         animals.Boar.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      sScrofa}
)
