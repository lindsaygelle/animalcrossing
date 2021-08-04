package koharu

import (
	"github.com/lindsaygelle/animalcrossing/animal/kangaroo"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "koharu"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Koharu.
	Villager = villager.Villager{
		Animal:      kangaroo.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Koharu,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
