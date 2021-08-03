package amelia

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "amelia"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Amelia.
	Villager = villager.Villager{
		Animal:  eagle.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
