package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	bam string = "Bam"
)

var (
	// Bam is an Animal Crossing villager.
	//
	// Bam is a Deer.
	Bam Villager = villager{
		animal: animals.Deer.Name(),
		name:   bam}
)
