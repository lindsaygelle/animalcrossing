package tipper

import (
	"github.com/lindsaygelle/animalcrossing/animal/cow"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cow01"
)

const (
	gender string = "female"
)

const (
	id string = "tipper"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tipper.
	Villager = villager.Villager{
		Animal:      cow.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tipper,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
