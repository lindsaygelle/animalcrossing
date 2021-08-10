package biff

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "hip04"
)

const (
	gender string = "male"
)

const (
	id string = "biff"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Biff.
	Villager = villager.Villager{
		Animal:      hippopotamus.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Biff,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
