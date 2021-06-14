package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Reindeer is the species information for Reindeers.
	Reindeer Species = species{
		class:        mammalia,
		conservation: vulnerable,
		domain:       eukarya,
		family:       cervidae,
		fictional:    (!fictional),
		genus:        rangifer,
		kingdom:      animalia,
		name:         animals.Reindeer.Name(),
		order:        artiodactyla,
		phylum:       chordata,
		species:      rTarandus}
)
