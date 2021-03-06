package stinky

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat13"
)

const (
	gender string = "male"
)

const (
	id string = "stinky"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Stinky.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Stinky,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
