package otis

import (
	"github.com/lindsaygelle/animalcrossing/animal/bird"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "otis"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Otis.
	Villager = villager.Villager{
		Animal:      bird.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Otis,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
