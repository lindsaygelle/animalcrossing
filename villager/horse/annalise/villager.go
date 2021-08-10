package annalise

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "hrs09"
)

const (
	gender string = "female"
)

const (
	id string = "annalise"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Annalise.
	Villager = villager.Villager{
		Animal:      horse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Annalise,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
