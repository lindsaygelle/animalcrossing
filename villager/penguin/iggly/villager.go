package iggly

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pgn11"
)

const (
	gender string = "male"
)

const (
	id string = "iggly"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Iggly.
	Villager = villager.Villager{
		Animal:      penguin.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Iggly,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
