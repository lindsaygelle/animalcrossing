package plucky

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "chn10"
)

const (
	gender string = "female"
)

const (
	id string = "plucky"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Plucky.
	Villager = villager.Villager{
		Animal:      chicken.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Plucky,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
