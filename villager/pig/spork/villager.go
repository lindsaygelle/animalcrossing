package spork

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "spork"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Spork.
	Villager = villager.Villager{
		Animal:  pig.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
