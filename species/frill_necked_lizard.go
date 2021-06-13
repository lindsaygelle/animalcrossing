package species

import "github.com/lindsaygelle/animalcrossing/animals"

var (
	// FrillNeckedLizard is the species information for Frill-Necked Lizards.
	FrillNeckedLizard Species = species{
		class:        sauropsida,
		conservation: leastConcern,
		domain:       eukarya,
		family:       agamidae,
		genus:        chlamydosaurus,
		kingdom:      animalia,
		name:         animals.FrillNeckedLizard.Name(),
		order:        squamata,
		phylum:       chordata,
		species:      cKingii}
)
