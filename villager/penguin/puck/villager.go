package puck

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pgn06"
)

const (
	gender string = "male"
)

const (
	id string = "puck"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Puck.
	Villager = villager.Villager{
		Animal:      penguin.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Puck,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
