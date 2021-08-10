package ketchup

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "duk13"
)

const (
	gender string = "female"
)

const (
	id string = "ketchup"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ketchup.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Ketchup,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
