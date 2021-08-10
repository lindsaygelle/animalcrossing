package frita

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "shp11"
)

const (
	gender string = "female"
)

const (
	id string = "frita"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Frita.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Frita,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
