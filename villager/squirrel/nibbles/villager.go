package nibbles

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ04"
)

const (
	gender string = "female"
)

const (
	id string = "nibbles"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Nibbles.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Nibbles,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
