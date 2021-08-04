package tangy

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "tangy"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Tangy.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Tangy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
