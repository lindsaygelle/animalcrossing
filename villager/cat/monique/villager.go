package monique

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat11"
)

const (
	gender string = "female"
)

const (
	id string = "monique"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Monique.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Monique,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
