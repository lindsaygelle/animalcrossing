package aisle

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "aisle"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Aisle.
	Villager = villager.Villager{
		Animal:  bearcub.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
