package grizzly

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "bea09"
)

const (
	gender string = "male"
)

const (
	id string = "grizzly"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Grizzly.
	Villager = villager.Villager{
		Animal:      bear.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Grizzly,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
