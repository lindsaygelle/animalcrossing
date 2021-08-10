package pekoe

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cbr14"
)

const (
	gender string = "female"
)

const (
	id string = "pekoe"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Pekoe.
	Villager = villager.Villager{
		Animal:      bearcub.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Pekoe,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
