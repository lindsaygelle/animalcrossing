package cyd

import (
	"github.com/lindsaygelle/animalcrossing/animal/elephant"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "male"
)

const (
	id string = "cyd"
)

const (
	personality string = "cranky"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cyd.
	Villager = villager.Villager{
		Animal:      elephant.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cyd,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
