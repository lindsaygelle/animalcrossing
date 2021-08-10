package cheri

import (
	"github.com/lindsaygelle/animalcrossing/animal/bearcub"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cbr10"
)

const (
	gender string = "female"
)

const (
	id string = "cheri"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Cheri.
	Villager = villager.Villager{
		Animal:      bearcub.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Cheri,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
