package peaches

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "hrs08"
)

const (
	gender string = "female"
)

const (
	id string = "peaches"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Peaches.
	Villager = villager.Villager{
		Animal:      horse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Peaches,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
