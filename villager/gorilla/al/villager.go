package al

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "gor08"
)

const (
	gender string = "male"
)

const (
	id string = "al"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Al.
	Villager = villager.Villager{
		Animal:      gorilla.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Al,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
