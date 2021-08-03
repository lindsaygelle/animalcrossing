package rhonda

import (
	"github.com/lindsaygelle/animalcrossing/animal/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "rhonda"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rhonda.
	Villager = villager.Villager{
		Animal:  rhinoceros.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
