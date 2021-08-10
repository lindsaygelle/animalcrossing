package ricky

import (
	"github.com/lindsaygelle/animalcrossing/animal/squirrel"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "squ10"
)

const (
	gender string = "male"
)

const (
	id string = "ricky"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ricky.
	Villager = villager.Villager{
		Animal:      squirrel.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Ricky,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
