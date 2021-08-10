package truffles

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pig01"
)

const (
	gender string = "female"
)

const (
	id string = "truffles"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Truffles.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Truffles,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
