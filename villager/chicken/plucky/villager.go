package plucky

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "plucky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Plucky.
	Villager = villager.Villager{
		Animal:  chicken.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
