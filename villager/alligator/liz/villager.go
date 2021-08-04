package liz

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "liz"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Liz.
	Villager = villager.Villager{
		Animal:      alligator.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Liz,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
