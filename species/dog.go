package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Dog is the species information for Dogs.
	Dog Species = species{
		class:        mammalia,
		conservation: domesticated,
		domain:       eukarya,
		family:       canidae,
		fictional:    (!fictional),
		genus:        canis,
		kingdom:      animalia,
		name:         animals.Dog.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      canisLupusFamiliaris}
)
