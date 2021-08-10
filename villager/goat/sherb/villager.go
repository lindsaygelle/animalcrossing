package sherb

import (
	"github.com/lindsaygelle/animalcrossing/animal/goat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "goa09"
)

const (
	gender string = "male"
)

const (
	id string = "sherb"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Sherb.
	Villager = villager.Villager{
		Animal:      goat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Sherb,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
