package biff

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "biff"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Biff.
	Villager = villager.Villager{
		Animal:  hippopotamus.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
