package tiara

import (
	"github.com/lindsaygelle/animalcrossing/animal/rhinoceros"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "tiara"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tiara.
	Villager = villager.Villager{
		Animal:      rhinoceros.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tiara,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
