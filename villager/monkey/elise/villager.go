package elise

import (
	"github.com/lindsaygelle/animalcrossing/animal/monkey"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mnk05"
)

const (
	gender string = "female"
)

const (
	id string = "elise"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Elise.
	Villager = villager.Villager{
		Animal:      monkey.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Elise,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
