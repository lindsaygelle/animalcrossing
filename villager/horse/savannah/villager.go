package savannah

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "hrs02"
)

const (
	gender string = "female"
)

const (
	id string = "savannah"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Savannah.
	Villager = villager.Villager{
		Animal:      horse.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Savannah,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
