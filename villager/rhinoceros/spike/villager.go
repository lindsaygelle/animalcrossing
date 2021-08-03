package spike

import (
	"github.com/lindsaygelle/animalcrossing/animal/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "spike"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Spike.
	Villager = villager.Villager{
		Animal:  rhinoceros.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
