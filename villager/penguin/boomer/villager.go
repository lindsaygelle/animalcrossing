package boomer

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pgn10"
)

const (
	gender string = "male"
)

const (
	id string = "boomer"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Boomer.
	Villager = villager.Villager{
		Animal:      penguin.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Boomer,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
