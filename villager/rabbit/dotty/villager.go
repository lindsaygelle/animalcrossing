package dotty

import (
	"github.com/lindsaygelle/animalcrossing/animal/rabbit"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "rbt01"
)

const (
	gender string = "female"
)

const (
	id string = "dotty"
)

const (
	personality string = "peppy"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Dotty.
	Villager = villager.Villager{
		Animal:      rabbit.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Dotty,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
