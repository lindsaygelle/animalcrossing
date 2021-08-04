package bangle

import (
	"github.com/lindsaygelle/animalcrossing/animal/tiger"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "bangle"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Bangle.
	Villager = villager.Villager{
		Animal:      tiger.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Bangle,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
