package punchy

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat06"
)

const (
	gender string = "male"
)

const (
	id string = "punchy"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Punchy.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Punchy,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
