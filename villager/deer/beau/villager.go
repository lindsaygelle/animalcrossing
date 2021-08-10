package beau

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "der07"
)

const (
	gender string = "male"
)

const (
	id string = "beau"
)

const (
	personality string = "lazy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Beau.
	Villager = villager.Villager{
		Animal:      deer.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Beau,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
