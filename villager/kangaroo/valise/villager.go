package valise

import (
	"github.com/lindsaygelle/animalcrossing/animal/kangaroo"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "valise"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Valise.
	Villager = villager.Villager{
		Animal:  kangaroo.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
