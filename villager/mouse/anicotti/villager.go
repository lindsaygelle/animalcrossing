package anicotti

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mus10"
)

const (
	gender string = "female"
)

const (
	id string = "anicotti"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Anicotti.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Anicotti,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
