package anicotti

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "anicotti"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Anicotti.
	Villager = villager.Villager{
		Animal:  mouse.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
