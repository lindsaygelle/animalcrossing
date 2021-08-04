package gwen

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "gwen"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gwen.
	Villager = villager.Villager{
		Animal:      penguin.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Gwen,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
