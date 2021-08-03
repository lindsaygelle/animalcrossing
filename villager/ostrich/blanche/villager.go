package blanche

import (
	"github.com/lindsaygelle/animalcrossing/animal/ostrich"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "blanche"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Blanche.
	Villager = villager.Villager{
		Animal:  ostrich.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
