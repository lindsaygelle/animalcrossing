package gala

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "gala"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gala.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Gala,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
