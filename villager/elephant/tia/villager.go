package tia

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "tia"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tia.
	Villager = villager.Villager{
		Animal:      elephant.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tia,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
