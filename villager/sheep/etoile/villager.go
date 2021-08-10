package etoile

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "shp14"
)

const (
	gender string = "female"
)

const (
	id string = "etoile"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Etoile.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Etoile,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
