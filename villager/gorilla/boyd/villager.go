package boyd

import (
	"github.com/lindsaygelle/animalcrossing/animal/gorilla"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "gor05"
)

const (
	gender string = "male"
)

const (
	id string = "boyd"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Boyd.
	Villager = villager.Villager{
		Animal:      gorilla.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Boyd,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
