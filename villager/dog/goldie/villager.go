package goldie

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "dog00"
)

const (
	gender string = "female"
)

const (
	id string = "goldie"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Goldie.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Goldie,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
