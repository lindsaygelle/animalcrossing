package mitzi

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat01"
)

const (
	gender string = "female"
)

const (
	id string = "mitzi"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Mitzi.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Mitzi,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
