package tbone

import (
	"github.com/lindsaygelle/animalcrossing/animal/bull"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "tbone"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for TBone.
	Villager = villager.Villager{
		Animal:  bull.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
