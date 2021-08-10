package lolly

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat18"
)

const (
	gender string = "female"
)

const (
	id string = "lolly"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Lolly.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Lolly,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
