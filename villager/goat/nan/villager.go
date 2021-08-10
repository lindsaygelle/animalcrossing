package nan

import (
	"github.com/lindsaygelle/animalcrossing/animal/goat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "goa01"
)

const (
	gender string = "female"
)

const (
	id string = "nan"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Nan.
	Villager = villager.Villager{
		Animal:      goat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Nan,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
