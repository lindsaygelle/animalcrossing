package fuchsia

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "der06"
)

const (
	gender string = "female"
)

const (
	id string = "fuchsia"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Fuchsia.
	Villager = villager.Villager{
		Animal:      deer.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Fuchsia,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
