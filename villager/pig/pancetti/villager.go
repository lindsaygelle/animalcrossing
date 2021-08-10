package pancetti

import (
	"github.com/lindsaygelle/animalcrossing/animal/pig"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "pig16"
)

const (
	gender string = "female"
)

const (
	id string = "pancetti"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Pancetti.
	Villager = villager.Villager{
		Animal:      pig.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Pancetti,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
