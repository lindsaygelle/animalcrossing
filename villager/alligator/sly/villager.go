package sly

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "sly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sly.
	Villager = villager.Villager{
		Animal:  alligator.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
