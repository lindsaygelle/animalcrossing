package quetzal

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "quetzal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Quetzal.
	Villager = villager.Villager{
		Animal:  eagle.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
