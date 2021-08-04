package reneigh

import (
	"github.com/lindsaygelle/animalcrossing/animal/horse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "reneigh"
)

const (
	personality string = "sisterly"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Reneigh.
	Villager = villager.Villager{
		Animal:      horse.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Reneigh,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
