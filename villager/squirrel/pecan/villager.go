package pecan

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ03"
)

const (
	gender string = "female"
)

const (
	id string = "pecan"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Pecan.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Pecan,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
