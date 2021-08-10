package rolf

import (
	"github.com/lindsaygelle/animalcrossing/animal/tiger"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "tig00"
)

const (
	gender string = "male"
)

const (
	id string = "rolf"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Rolf.
	Villager = villager.Villager{
		Animal:      tiger.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Rolf,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
