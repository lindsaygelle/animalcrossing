package kidd

import (
	"github.com/lindsaygelle/animalcrossing/animal/goat"
	"github.com/lindsaygelle/animalcrossing/villager"
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
		Gender:      gender,
		Id:          id,
		Key:         villager.Kidd,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
