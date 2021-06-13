package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// Axolotl is the species information for Axolotls.
	Axolotl Species = species{
		class:        amphibia,
		conservation: criticallyEndangered,
		domain:       eukarya,
		family:       ambystomatidae,
		genus:        ambystoma,
		kingdom:      animalia,
		name:         animals.Axolotl.Name(),
		order:        caudata,
		phylum:       chordata,
		species:      aMexicanum}
)
