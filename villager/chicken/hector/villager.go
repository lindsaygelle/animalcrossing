package hector

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "hector"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Hector.
	Villager = villager.Villager{
		Animal:      chicken.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Hector,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
