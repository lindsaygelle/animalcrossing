package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	Cat Species = species{
		class:        mammalia,
		conservation: domesticated,
		domain:       eukarya,
		family:       felidae,
		genus:        felis,
		kingdom:      animalia,
		name:         animals.Cat.Name(),
		order:        carnivora,
		phylum:       chordata,
		species:      fCatus}
)
