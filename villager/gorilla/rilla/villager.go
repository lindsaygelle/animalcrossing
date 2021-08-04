package rilla

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "rilla"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rilla.
	Villager = villager.Villager{
		Animal:      gorilla.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rilla,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
