package goose

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "goose"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Goose.
	Villager = villager.Villager{
		Animal:      chicken.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Goose,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
