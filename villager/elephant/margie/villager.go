package margie

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "elp04"
)

const (
	gender string = "female"
)

const (
	id string = "margie"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Margie.
	Villager = villager.Villager{
		Animal:      elephant.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Margie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
