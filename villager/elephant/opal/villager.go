package opal

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "opal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Opal.
	Villager = villager.Villager{
		Animal:  elephant.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
