package champ

import (
	"github.com/lindsaygelle/animalcrossing/animal/monkey"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mnk00"
)

const (
	gender string = "male"
)

const (
	id string = "champ"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Champ.
	Villager = villager.Villager{
		Animal:      monkey.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Champ,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
