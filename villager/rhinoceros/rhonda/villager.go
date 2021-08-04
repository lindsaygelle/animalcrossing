package rhonda

import (
	"github.com/lindsaygelle/animalcrossing/animal/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "rhonda"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rhonda.
	Villager = villager.Villager{
		Animal:      rhinoceros.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rhonda,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
