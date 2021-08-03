package hippeux

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "hippeux"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Hippeux.
	Villager = villager.Villager{
		Animal:  hippopotamus.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
