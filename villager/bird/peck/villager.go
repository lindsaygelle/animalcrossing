package peck

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "peck"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Peck.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Peck,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
