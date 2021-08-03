package rowan

import (
	"github.com/lindsaygelle/animalcrossing/animal/tiger"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "rowan"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rowan.
	Villager = villager.Villager{
		Animal:  tiger.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
