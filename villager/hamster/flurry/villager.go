package flurry

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "flurry"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Flurry.
	Villager = villager.Villager{
		Animal:      hamster.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Flurry,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
