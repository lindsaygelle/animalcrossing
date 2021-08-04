package tad

import (
	"github.com/lindsaygelle/animalcrossing/animal/frog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "tad"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tad.
	Villager = villager.Villager{
		Animal:      frog.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tad,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
