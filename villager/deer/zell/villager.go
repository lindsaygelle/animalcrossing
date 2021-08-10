package zell

import (
	"github.com/lindsaygelle/animalcrossing/animal/deer"
	"github.com/lindsaygelle/animalcrossing/villager"
)

const (
	code string = "der02"
)

const (
	gender string = "male"
)

const (
	id string = "zell"
)

const (
	personality string = "smug"
)

const (
	special bool = false
)

var (
	// Villager is the villager information for Zell.
	Villager = villager.Villager{
		Animal:      deer.Animal,
		Birthday:    birthday,
		Gender:      gender,
		Id:          id,
		Key:         villager.Zell,
		Name:        name,
		Personality: personality,
		Phrase:      phrase,
		Special:     special}
)
