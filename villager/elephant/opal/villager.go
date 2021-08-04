package opal

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "opal"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Opal.
	Villager = villager.Villager{
		Animal:      elephant.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Opal,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
