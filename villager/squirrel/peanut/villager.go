package peanut

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ00"
)

const (
	gender string = "female"
)

const (
	id string = "peanut"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Peanut.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Peanut,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
