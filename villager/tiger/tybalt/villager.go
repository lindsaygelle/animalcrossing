package tybalt

import (
	"github.com/lindsaygelle/animalcrossing/animal/tiger"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "tybalt"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tybalt.
	Villager = villager.Villager{
		Animal:  tiger.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
