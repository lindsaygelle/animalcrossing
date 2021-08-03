package flurry

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	id string = "flurry"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Flurry.
	Villager = villager.Villager{
		Animal:  hamster.Animal,
		Id:      id,
		Name:    name,
		Phrase:  phrase,
		Special: special}
)
