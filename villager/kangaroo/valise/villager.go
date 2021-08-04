package valise

import (
	"github.com/lindsaygelle/animalcrossing/animal/kangaroo"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "valise"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Valise.
	Villager = villager.Villager{
		Animal:      kangaroo.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Valise,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
