package alli

import (
	"github.com/lindsaygelle/animalcrossing/animal/alligator"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "crd01"
)

const (
	gender string = "female"
)

const (
	id string = "alli"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Alli.
	Villager = villager.Villager{
		Animal:      alligator.Animal,
		Birthday:    birthday,
		Code:        code,
		Gender:      gender,
		Id:          id,
		Key:         villager.Alli,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
