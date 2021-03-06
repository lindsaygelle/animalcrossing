package scoot

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "duk10"
)

const (
	gender string = "male"
)

const (
	id string = "scoot"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Scoot.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Scoot,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
