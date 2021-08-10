package miranda

import (
	"github.com/lindsaygelle/animalcrossing/animal/duck"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "duk12"
)

const (
	gender string = "female"
)

const (
	id string = "miranda"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Miranda.
	Villager = villager.Villager{
		Animal:      duck.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Miranda,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
