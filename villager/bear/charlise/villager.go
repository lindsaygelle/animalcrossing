package charlise

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "charlise"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Charlise.
	Villager = villager.Villager{
		Animal:  bear.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
