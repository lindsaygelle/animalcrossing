package maelle

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "maelle"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Maelle.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Maelle,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
