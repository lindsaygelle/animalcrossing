package cleo

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "cleo"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cleo.
	Villager = villager.Villager{
		Animal:      horse.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cleo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
