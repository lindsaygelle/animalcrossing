package dizzy

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "elp01"
)

const (
	gender string = "male"
)

const (
	id string = "dizzy"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Dizzy.
	Villager = villager.Villager{
		Animal:      elephant.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Dizzy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
