package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Monkey is the species information for Monkeys.
	Monkey Species = species{
		class:        mammalia,
		conservation: endangered,
		domain:       eukarya,
		family:       na,
		fictional:    (!fictional),
		genus:        na,
		kingdom:      animalia,
		name:         animals.Monkey.Name(),
		order:        primates,
		phylum:       chordata,
		species:      na}
)
