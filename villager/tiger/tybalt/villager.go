package tybalt

import (
	"github.com/lindsaygelle/animalcrossing/animal/tiger"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "tig02"
)

const (
	gender string = "male"
)

const (
	id string = "tybalt"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tybalt.
	Villager = villager.Villager{
		Animal:      tiger.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tybalt,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
