package chadder

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mus18"
)

const (
	gender string = "male"
)

const (
	id string = "chadder"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Chadder.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Chadder,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
