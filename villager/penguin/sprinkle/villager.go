package sprinkle

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pgn14"
)

const (
	gender string = "female"
)

const (
	id string = "sprinkle"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sprinkle.
	Villager = villager.Villager{
		Animal:      penguin.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Sprinkle,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
