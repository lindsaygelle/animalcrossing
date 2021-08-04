package nana

import (
	"github.com/lindsaygelle/animalcrossing/animal/monkey"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "nana"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Nana.
	Villager = villager.Villager{
		Animal:      monkey.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Nana,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
