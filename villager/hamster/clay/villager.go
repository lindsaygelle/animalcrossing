package clay

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "clay"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Clay.
	Villager = villager.Villager{
		Animal:      hamster.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Clay,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
