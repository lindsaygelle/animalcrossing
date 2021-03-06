package rosie

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat02"
)

const (
	gender string = "female"
)

const (
	id string = "rosie"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rosie.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rosie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
