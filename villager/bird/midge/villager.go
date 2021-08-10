package midge

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "brd08"
)

const (
	gender string = "female"
)

const (
	id string = "midge"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Midge.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Midge,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
