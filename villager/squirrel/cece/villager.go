package cece

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ19"
)

const (
	gender string = "female"
)

const (
	id string = "cece"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cece.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cece,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
