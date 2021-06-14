package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Tapir is the species information for Tapirs.
	Tapir Species = species{
		class:        mammalia,
		conservation: unknown,
		domain:       eukarya,
		family:       tapiridae,
		genus:        tapirus,
		kingdom:      animalia,
		name:         animals.Tapir.Name(),
		order:        perissodactyla,
		phylum:       chordata,
		species:      na}
)
