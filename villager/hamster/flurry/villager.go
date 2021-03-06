package flurry

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ham06"
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
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Flurry,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
