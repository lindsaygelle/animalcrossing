package deli

import (
	"github.com/lindsaygelle/animalcrossing/animal/monkey"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "deli"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Deli.
	Villager = villager.Villager{
		Animal:  monkey.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
