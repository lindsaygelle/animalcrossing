package chrissy

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "chrissy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Chrissy.
	Villager = villager.Villager{
		Animal:  rabbit.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
