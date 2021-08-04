package ava

import (
	"github.com/lindsaygelle/animalcrossing/animal/chicken"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "ava"
)

const (
	personality string = "normal"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Ava.
	Villager = villager.Villager{
		Animal:      chicken.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Ava,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
