package violet

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "gor07"
)

const (
	gender string = "female"
)

const (
	id string = "violet"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Violet.
	Villager = villager.Villager{
		Animal:      gorilla.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Violet,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
