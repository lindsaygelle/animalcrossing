package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Turkey is the species information for Turkeys.
	Turkey Species = species{
		class:        aves,
		conservation: unknown,
		domain:       eukarya,
		family:       phasianidae,
		fictional:    (!fictional),
		genus:        meleagris,
		kingdom:      animalia,
		name:         animals.Turkey.Name(),
		order:        galliformes,
		phylum:       chordata,
		species:      na,
	}
)
