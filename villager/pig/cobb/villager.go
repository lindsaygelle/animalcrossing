package cobb

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pig08"
)

const (
	gender string = "male"
)

const (
	id string = "cobb"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cobb.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cobb,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
