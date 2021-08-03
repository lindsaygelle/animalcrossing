package octavian

import (
	"github.com/lindsaygelle/animalcrossing/animal/octopus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "octavian"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Octavian.
	Villager = villager.Villager{
		Animal:  octopus.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
