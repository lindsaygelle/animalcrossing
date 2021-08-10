package kitty

import (
	"github.com/lindsaygelle/animalcrossing/animal/cat"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "cat14"
)

const (
	gender string = "female"
)

const (
	id string = "kitty"
)

const (
	personality string = "snooty"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Kitty.
	Villager = villager.Villager{
		Animal:      cat.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Kitty,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
