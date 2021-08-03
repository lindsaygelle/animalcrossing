package peaches

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "peaches"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Peaches.
	Villager = villager.Villager{
		Animal:  horse.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
