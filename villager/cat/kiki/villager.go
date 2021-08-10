package kiki

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat04"
)

const (
	gender string = "female"
)

const (
	id string = "kiki"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Kiki.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Kiki,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
