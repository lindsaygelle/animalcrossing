package sly

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "crd06"
)

const (
	gender string = "male"
)

const (
	id string = "sly"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sly.
	Villager = villager.Villager{
		Animal:      alligator.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Sly,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
