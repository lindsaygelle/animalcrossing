package tutu

import (
	"github.com/lindsaygelle/animalcrossing/animal/bear"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "bea07"
)

const (
	gender string = "female"
)

const (
	id string = "tutu"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tutu.
	Villager = villager.Villager{
		Animal:      bear.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tutu,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
