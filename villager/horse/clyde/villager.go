package clyde

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "hrs10"
)

const (
	gender string = "male"
)

const (
	id string = "clyde"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Clyde.
	Villager = villager.Villager{
		Animal:      horse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Clyde,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
