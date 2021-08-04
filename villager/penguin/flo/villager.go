package flo

import (
	"github.com/lindsaygelle/animalcrossing/animal/penguin"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "flo"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Flo.
	Villager = villager.Villager{
		Animal:      penguin.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Flo,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
