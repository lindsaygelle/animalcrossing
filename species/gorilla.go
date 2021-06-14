package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Gorrilla is the species information for Gorillas.
	Gorilla Species = species{
		class:        mammalia,
		conservation: endangered,
		domain:       eukarya,
		family:       hominidae,
		fictional:    (!fictional),
		genus:        gorilla,
		kingdom:      animalia,
		name:         animals.Gorilla.Name(),
		order:        primates,
		phylum:       chordata,
		species:      na}
)
