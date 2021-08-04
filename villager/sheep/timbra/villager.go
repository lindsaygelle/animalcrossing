package timbra

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "timbra"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Timbra.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Timbra,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
