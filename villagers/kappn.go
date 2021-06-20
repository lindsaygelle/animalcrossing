package villagers

import "github.com/lindsaygelle/animalcrossing/animals"

const (
	kappn string = "Kapp'n"
)

var (
	// Kappn is an Animal Crossing villager.
	//
	// Kappn is a Kappa.
	Kappn Villager = villager{
		animal: animals.Kappa.Name(),
		name:   kappn}
)
