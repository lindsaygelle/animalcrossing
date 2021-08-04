package gen

import (
	"github.com/lindsaygelle/animalcrossing/animal/sheep"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "gen"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Gen.
	Villager = villager.Villager{
		Animal:      sheep.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Gen,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
