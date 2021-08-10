package eloise

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "elp03"
)

const (
	gender string = "female"
)

const (
	id string = "eloise"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Eloise.
	Villager = villager.Villager{
		Animal:      elephant.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Eloise,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
