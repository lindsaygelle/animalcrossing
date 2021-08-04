package candi

import (
	"github.com/lindsaygelle/animalcrossing/animal/mouse"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "candi"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Candi.
	Villager = villager.Villager{
		Animal:      mouse.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Candi,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
