package iggly

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "iggly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Iggly.
	Villager = villager.Villager{
		Animal:  penguin.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
