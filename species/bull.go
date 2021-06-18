package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Bull is the species information for Bulls (Bovines).
	Bull Species = species{
		class:        Bovine.Class(),
		conservation: Bovine.Conservation(),
		domain:       Bovine.Domain(),
		family:       Bovine.Family(),
		fictional:    Bovine.Fictional(),
		genus:        Bovine.Genus(),
		kingdom:      Bovine.Kingdom(),
		name:         animals.Bull.Name(),
		order:        Bovine.Order(),
		phylum:       Bovine.Phylum(),
		species:      Bovine.Species()}
)
