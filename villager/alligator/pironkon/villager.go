package pironkon

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "pironkon"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Pironkon.
	Villager = villager.Villager{
		Animal:      alligator.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Pironkon,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
