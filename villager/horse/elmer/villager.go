package elmer

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "elmer"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Elmer.
	Villager = villager.Villager{
		Animal:      horse.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Elmer,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
