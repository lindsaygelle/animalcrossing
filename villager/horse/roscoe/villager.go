package roscoe

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "hrs04"
)

const (
	gender string = "male"
)

const (
	id string = "roscoe"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Roscoe.
	Villager = villager.Villager{
		Animal:      horse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Roscoe,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
