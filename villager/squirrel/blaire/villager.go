package blaire

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ01"
)

const (
	gender string = "female"
)

const (
	id string = "blaire"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Blaire.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Blaire,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
