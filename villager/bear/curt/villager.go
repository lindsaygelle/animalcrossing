package curt

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "curt"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Curt.
	Villager = villager.Villager{
		Animal:  bear.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
