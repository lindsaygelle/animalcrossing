package megan

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "bea15"
)

const (
	gender string = "female"
)

const (
	id string = "megan"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Megan.
	Villager = villager.Villager{
		Animal:      bear.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Megan,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
