package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Chicken is the species information for Chickens.
	Chicken Species = species{
		class:        aves,
		conservation: domesticated,
		domain:       eukarya,
		family:       phasianidae,
		fictional:    (!fictional),
		genus:        gallus,
		kingdom:      animalia,
		name:         animals.Chicken.Name(),
		order:        galliformes,
		phylum:       chordata,
		species:      gGallus}
)
