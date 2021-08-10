package soleil

import (
	"github.com/lindsaygelle/animalcrossing/animal/hamster"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "ham04"
)

const (
	gender string = "female"
)

const (
	id string = "soleil"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Soleil.
	Villager = villager.Villager{
		Animal:      hamster.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Soleil,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
