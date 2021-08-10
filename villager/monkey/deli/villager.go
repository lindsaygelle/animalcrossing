package deli

import (
	"github.com/lindsaygelle/animalcrossing/animal/monkey"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "mnk08"
)

const (
	gender string = "male"
)

const (
	id string = "deli"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Deli.
	Villager = villager.Villager{
		Animal:      monkey.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Deli,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
