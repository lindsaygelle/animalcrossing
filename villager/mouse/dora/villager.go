package dora

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "dora"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Dora.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Dora,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
