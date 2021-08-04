package friga

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "friga"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Friga.
	Villager = villager.Villager{
		Animal:      penguin.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Friga,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
