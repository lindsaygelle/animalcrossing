package champ

import (
	"github.com/lindsaygelle/animalcrossing/animal/monkey"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "champ"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Champ.
	Villager = villager.Villager{
		Animal:  monkey.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
