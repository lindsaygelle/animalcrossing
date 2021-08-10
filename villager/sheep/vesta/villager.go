package vesta

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "shp00"
)

const (
	gender string = "female"
)

const (
	id string = "vesta"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Vesta.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Vesta,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
