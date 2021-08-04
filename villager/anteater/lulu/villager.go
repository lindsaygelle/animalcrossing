package lulu

import (
	"github.com/lindsaygelle/animalcrossing/animal/anteater"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "lulu"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Lulu.
	Villager = villager.Villager{
		Animal:      anteater.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Lulu,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
