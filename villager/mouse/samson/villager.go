package samson

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mus04"
)

const (
	gender string = "male"
)

const (
	id string = "samson"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Samson.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Samson,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
