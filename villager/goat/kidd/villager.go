package kidd

import (
	"github.com/lindsaygelle/animalcrossing/animal/goat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "goa07"
)

const (
	gender string = "male"
)

const (
	id string = "kidd"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Kidd.
	Villager = villager.Villager{
		Animal:      goat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Kidd,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
