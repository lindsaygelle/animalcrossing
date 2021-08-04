package lulu

import (
	"github.com/lindsaygelle/animalcrossing/animal/hippopotamus"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	gender string = "female"
)

const (
	id string = "lulu"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Lulu.
	Villager = villager.Villager{
		Animal:      hippopotamus.Animal,
		Gender:      gender,
		Id:          id,
		Key:         villager.Lulu,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
