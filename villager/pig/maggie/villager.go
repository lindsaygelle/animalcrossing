package maggie

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pig10"
)

const (
	gender string = "female"
)

const (
	id string = "maggie"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Maggie.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Maggie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
