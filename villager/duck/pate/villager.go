package pate

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "duk02"
)

const (
	gender string = "female"
)

const (
	id string = "pate"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Pate.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Pate,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
