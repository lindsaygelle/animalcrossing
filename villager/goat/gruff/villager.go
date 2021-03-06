package gruff

import (
	"github.com/lindsaygelle/animalcrossing/animal/goat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "goa04"
)

const (
	gender string = "male"
)

const (
	id string = "gruff"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gruff.
	Villager = villager.Villager{
		Animal:      goat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Gruff,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
