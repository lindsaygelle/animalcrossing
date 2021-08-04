package holden

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "holden"
)

const (
	personality string = "jock"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Holden.
	Villager = villager.Villager{
		Animal:      hamster.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Holden,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
