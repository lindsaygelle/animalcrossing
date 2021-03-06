package louie

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "gor04"
)

const (
	gender string = "male"
)

const (
	id string = "louie"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Louie.
	Villager = villager.Villager{
		Animal:      gorilla.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Louie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
