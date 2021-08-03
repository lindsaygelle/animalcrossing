package zucker

import (
	"github.com/lindsaygelle/animalcrossing/animal/octopus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "zucker"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Zucker.
	Villager = villager.Villager{
		Animal:  octopus.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
