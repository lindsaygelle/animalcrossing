package peewee

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "gor01"
)

const (
	gender string = "male"
)

const (
	id string = "peewee"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Peewee.
	Villager = villager.Villager{
		Animal:      gorilla.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Peewee,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
