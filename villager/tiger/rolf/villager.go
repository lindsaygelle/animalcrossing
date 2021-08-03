package rolf

import (
	"github.com/lindsaygelle/animalcrossing/animal/tiger"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "rolf"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rolf.
	Villager = villager.Villager{
		Animal:  tiger.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
