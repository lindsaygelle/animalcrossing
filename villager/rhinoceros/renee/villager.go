package renee

import (
	"github.com/lindsaygelle/animalcrossing/animal/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "renee"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Renee.
	Villager = villager.Villager{
		Animal:  rhinoceros.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
