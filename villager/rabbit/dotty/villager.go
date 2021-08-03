package dotty

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "dotty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Dotty.
	Villager = villager.Villager{
		Animal:  rabbit.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
