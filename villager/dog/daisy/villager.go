package daisy

import (
	"github.com/lindsaygelle/animalcrossing/animal/dog"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "dog07"
)

const (
	gender string = "female"
)

const (
	id string = "daisy"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Daisy.
	Villager = villager.Villager{
		Animal:      dog.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Daisy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
