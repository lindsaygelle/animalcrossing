package freckles

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "duk07"
)

const (
	gender string = "female"
)

const (
	id string = "freckles"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Freckles.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Freckles,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
