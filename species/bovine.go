package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Bovine is the species information for Bulls & Cows.
	Bovine Species = species{
		class:        mammalia,
		conservation: domesticated,
		domain:       eukarya,
		family:       bovidae,
		genus:        bos,
		kingdom:      animalia,
		name:         animals.Bovine.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      bTaurus}
)
