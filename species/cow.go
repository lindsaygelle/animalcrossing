package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Cow is the species information for Cows (Bovines).
	Cow Species = species{
		class:        Bovine.Class(),
		conservation: Bovine.Conservation(),
		domain:       Bovine.Domain(),
		family:       Bovine.Family(),
		fictional:    Bovine.Fictional(),
		genus:        Bovine.Genus(),
		kingdom:      Bovine.Kingdom(),
		name:         animals.Cow.Name(),
		order:        Bovine.Order(),
		phylum:       Bovine.Phylum(),
		species:      Bovine.Species()}
)
