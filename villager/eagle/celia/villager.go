package celia

import (
	"github.com/lindsaygelle/animalcrossing/animal/eagle"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "celia"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Celia.
	Villager = villager.Villager{
		Animal:      eagle.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Celia,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
