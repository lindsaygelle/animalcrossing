package drago

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "drago"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Drago.
	Villager = villager.Villager{
		Animal:  alligator.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
