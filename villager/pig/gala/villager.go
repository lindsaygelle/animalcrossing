package gala

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "gala"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gala.
	Villager = villager.Villager{
		Animal:  pig.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
